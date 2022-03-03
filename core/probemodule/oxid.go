package probemodule

import (
	"bytes"
	"cube/config"
	"cube/pkg"
	"fmt"
	"net"
	"strings"
)

type Oxid struct {
	*Probe
}

func (o Oxid) ProbeName() string {
	return "oxid"
}

func (o Oxid) ProbePort() string {
	return "135"
}

func (o Oxid) ProbeLoad() bool {
	return true
}

func (o Oxid) PortCheck() bool {
	return true
}

func (o Oxid) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *o.Probe, Result: "", Err: nil}
	dl := net.Dialer{Timeout: config.TcpConnTimeout}
	t := fmt.Sprintf("%s:%s", o.Ip, o.Port)
	conn, err := dl.Dial("tcp", t)

	// defer conn.Close()
	if err != nil {
		result.Err = err
		//log.Printf("Oxid Running Error: %s:%s", task.Ip, err)
		return result
	}

	conn.Write([]byte("\x05\x00\x0b\x03\x10\x00\x00\x00\x48\x00\x00\x00\x01\x00\x00\x00\xb8\x10\xb8\x10\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x01\x00\xc4\xfe\xfc\x99\x60\x52\x1b\x10\xbb\xcb\x00\xaa\x00\x21\x34\x7a\x00\x00\x00\x00\x04\x5d\x88\x8a\xeb\x1c\xc9\x11\x9f\xe8\x08\x00\x2b\x10\x48\x60\x02\x00\x00\x00"))

	tmpByte := make([]byte, 1024)
	conn.Read(tmpByte)

	// dcerpc finish

	// IOXIDResolve start
	conn.Write([]byte("\x05\x00\x00\x03\x10\x00\x00\x00\x18\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00"))

	r := make([]byte, 4096)
	_, err = conn.Read(r)
	//if err != nil {
	//	return err
	//}
	/*
		Two parts:
			1. Distributed Computing Enviroment / Remote Prodedure Call Response
			2. DCOM OXID Resolver <- what we need
	*/

	r = r[24+12+2+2:]
	index := bytes.Index(r, []byte("\x09\x00\xff\xff\x00\x00"))
	//if index == -1 {
	//	return errors.New("Not Found")
	//}
	r = r[:index]
	var results []string

	for {
		if len(r) == 0 {
			break
		}
		index = bytes.Index(r, []byte("\x00\x00\x00"))
		hosts := pkg.Bytes2StringUTF16(r[:index+3])
		results = append(results, hosts)
		r = r[index+3:]
	}
	//var hostname string
	//re := regexp.MustCompile("[[:^ascii:]]")

	var netAddr []string
	if len(results) > 0 {
		//hostname := re.ReplaceAllLiteralString(results[0], "")
		hostname := results[0]
		for _, v := range results[1:] {
			netAddr = append(netAddr, v)
		}
		result.Result = fmt.Sprintf("Host: %s\nNets: %s\n", hostname, strings.Join(netAddr, "\t"))
	}
	arch := getArch(o.Ip, o.Port)
	if len(arch) > 0 {
		result.Result += fmt.Sprintf("Arch: %s\n", arch)
	}
	return result
}

func init() {
	AddProbeKeys("oxid")
}

func getArch(ip, port string) (s string) {
	dl := net.Dialer{Timeout: config.TcpConnTimeout}
	t := fmt.Sprintf("%s:%s", ip, port)
	conn, err := dl.Dial("tcp", t)
	if err != nil {
		return
	}

	archPayload := []byte{ /* Packet 186 */
		0x05, 0x00, 0x0b, 0x03, 0x10, 0x00, 0x00, 0x00,
		0x48, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0xb8, 0x10, 0xb8, 0x10, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00,
		0x08, 0x83, 0xaf, 0xe1, 0x1f, 0x5d, 0xc9, 0x11,
		0x91, 0xa4, 0x08, 0x00, 0x2b, 0x14, 0xa0, 0xfa,
		0x03, 0x00, 0x00, 0x00, 0x33, 0x05, 0x71, 0x71,
		0xba, 0xbe, 0x37, 0x49, 0x83, 0x19, 0xb5, 0xdb,
		0xef, 0x9c, 0xcc, 0x36, 0x01, 0x00, 0x00, 0x00,
	}
	conn.Write(archPayload)
	tmpByte := make([]byte, 60)
	conn.Read(tmpByte)
	if bytes.Index(tmpByte, []byte("\x33\x05\x71\x71\xba\xbe\x37\x49\x83\x19\xb5\xdb\xef\x9c\xcc\x36")) > 0 {
		s = "64-bit"
	}
	if strings.Contains(string(tmpByte), "syntaxes_not_supported") {
		s = "32-bit"
	}
	return s
}

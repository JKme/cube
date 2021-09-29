package probe

import (
	"bytes"
	"cube/model"
	"cube/util"
	"fmt"
	"net"
	"strings"
)

func OxidProbe(task model.ProbeTask) (result model.ProbeTaskResult) {
	result = model.ProbeTaskResult{ProbeTask: task, Result: "", Err: nil}
	//log.Printf("Oxid Running Debug: %s", task.Ip)
	dl := net.Dialer{Timeout: model.ConnectTimeout}
	t := fmt.Sprintf("%s:%s", task.Ip, task.Port)
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
		hosts := util.Bytes2StringUTF16(r[:index+3])
		results = append(results, hosts)
		r = r[index+3:]
	}
	//var hostname string
	var netAddr []string
	if len(results) > 0 {
		hostname := results[0]
		for _, v := range results[1:] {
			netAddr = append(netAddr, v)
		}
		result.Result = fmt.Sprintf("Host: %s\nNets: %s\n", hostname, strings.Join(netAddr, "\t"))
	}
	arch := getArch(task)
	if len(arch) > 0 {
		result.Result += fmt.Sprintf("Arch: %s\n", arch)
	}
	//edr := getEDR(task)
	//fmt.Println(edr)
	//if len(edr) > 0 {
	//	result.Result += fmt.Sprintf("EDR: %s\n", edr)
	//}
	return result
}

func getArch(task model.ProbeTask) (s string) {
	dl := net.Dialer{Timeout: model.ConnectTimeout}
	t := fmt.Sprintf("%s:%s", task.Ip, task.Port)
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

//func getEDR(task model.ProbeTask) (s string) {
//	dl := net.Dialer{Timeout: model.ConnectTimeout}
//	t := fmt.Sprintf("%s:%s", task.Ip, task.Port)
//	conn, err := dl.Dial("tcp", t)
//	if err != nil {
//		return
//	}
//
//	p1 := []byte{ /* Packet 186 */
//		0x05, 0x00, 0x0b, 0x03, 0x10, 0x00, 0x00, 0x00,
//		0x48, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
//		0xb8, 0x10, 0xb8, 0x10, 0x00, 0x00, 0x00, 0x00,
//		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00,
//		0x08, 0x83, 0xaf, 0xe1, 0x1f, 0x5d, 0xc9, 0x11,
//		0x91, 0xa4, 0x08, 0x00, 0x2b, 0x14, 0xa0, 0xfa,
//		0x03, 0x00, 0x00, 0x00, 0x33, 0x05, 0x71, 0x71,
//		0xba, 0xbe, 0x37, 0x49, 0x83, 0x19, 0xb5, 0xdb,
//		0xef, 0x9c, 0xcc, 0x36, 0x01, 0x00, 0x00, 0x00,
//	}
//	conn.Write(p1)
//
//	r := make([]byte, 1024)
//	conn.Read(r)
//
//
//	edrPayload := []byte(
//		"\x00\x0c\x29\x6b\x33\xf5\x0a\xf8\xbc\x36\x72\x64\x08\x00\x45\x00" +
//		"\x00\x68\x00\x00\x00\x00\x40\x06\x00\x00\xac\x10\x9d\x01\xac\x10" +
//		"\x9d\x04\xec\xf1\x00\x87\x48\x33\x55\x45\xb0\x3f\x5a\x5a\x50\x18" +
//		"\x10\x00\x92\x81\x00\x00\x05\x00\x00\x03\x10\x00\x00\x00\x40\x00" +
//		"\x00\x00\x00\x00\x00\x00\x28\x00\x00\x00\x00\x00\x02\x00\x00\x00" +
//		"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
//		"\x00\x00\x88\xc9\x50\xf0\x54\x48\x9b\x4c\x99\xe0\x64\x72\xc9\x76" +
//		"\xf9\xaa\x01\x00\x00\x00",
//	)
//	conn.Write(edrPayload)
//	tmpByte := make([]byte, 60)
//	conn.Read(tmpByte)
//	fmt.Println(string(tmpByte))
//	//if bytes.Index(tmpByte, []byte("\x33\x05\x71\x71\xba\xbe\x37\x49\x83\x19\xb5\xdb\xef\x9c\xcc\x36")) > 0 {
//	//	s = "AVG or AVAST"
//	//}
//	return s
//}

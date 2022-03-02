package probemodule

import (
	"bytes"
	"cube/config"
	"cube/pkg"
	"encoding/binary"
	"fmt"
	"github.com/JKme/go-ntlmssp"
	manuf "github.com/JKme/gomanuf"
	"math/rand"
	"net"
	"strings"
	"time"
)

type Netbios struct {
	*Probe
}

func (n Netbios) ProbeName() string {
	return "netbios"
}

func (n Netbios) ProbePort() string {
	return "137"
}

func (n Netbios) ProbeSkipPortCheck() bool {
	return true
}

func (n Netbios) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *n.Probe, Result: "", Err: nil}

	conn, err := net.DialTimeout("udp", fmt.Sprintf("%s:%v", n.Ip, n.Port), config.TcpConnTimeout)
	if err != nil {
		return result
	}
	defer conn.Close()

	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		//log.Println("SetReadDeadline failed:", err)
		return result
	}
	_, err = conn.Write(createStatusRequest())
	if err != nil {
		return result
	}
	ret2, err := pkg.ReadBytes(conn)
	if err != nil {
		return result
	}
	//fmt.Println(text)
	sreply := parseReplay(ret2)
	if len(sreply.Names) == 0 && len(sreply.Addresses) == 0 {
		return result
	}

	var nreply NetbiosReplyStatus

	if sreply.Header.RecordType == 0x21 {
		_, err = conn.Write(createNameRequest(pkg.TrimName(string(sreply.HostName[:]))))
		ret2, _ = pkg.ReadBytes(conn)
		nreply = parseReplay(ret2)
	}

	//if sreply.Header.RecordType == 0x00 {
	//	return
	//}

	//fmt.Printf("%x\n", sreply.Header.RecordType)
	//if sreply.Header.RecordType != 0x20 {
	//	return
	//}

	var Info map[string]string
	Info = make(map[string]string)

	//N, _ := util.GbkToUtf8(sreply.HostName[:])
	//fmt.Println(util.IsUtf8(sreply.HostName[:]))
	//fmt.Printf("%x\n", sreply.HostName[:])
	//fmt.Printf("%x\n", N)
	//fmt.Println(string(N))
	uniqueName = sreply.HostName[:]
	payload := getNbssPayload(uniqueName)
	ntlminfo := probe139(n, payload)

	Name, _ := pkg.ByteToString(sreply.HostName[:])
	if len(Name) > 0 {
		Info["Name"] = Name
	}
	var Nets []string
	if nreply.Header.RecordType == 0x20 {
		for _, ainfo := range nreply.Addresses {

			net1 := fmt.Sprintf("%d.%d.%d.%d", ainfo.Address[0], ainfo.Address[1], ainfo.Address[2], ainfo.Address[3])
			if net1 == "0.0.0.0" {
				continue
			}

			Nets = append(Nets, net1)
		}
	}

	if sreply.HWAddr != "00:00:00:00:00:00" {
		m1 := manuf.Search(sreply.HWAddr)
		Info["Hwaddr"] = sreply.HWAddr + fmt.Sprintf("\t\t%s", m1)
	}

	username := pkg.TrimName(string(sreply.UserName[:]))
	if len(username) > 0 && username != Info["Name"] {
		Info["Username"] = username
	}

	for _, rname := range sreply.Names {

		//tname := util.TrimName(string(rname.Name[:]))
		tname, _ := pkg.ByteToString(rname.Name[:])
		if tname == Info["Name"] {
			continue
		}

		if rname.Flag&0x0800 != 0 {
			continue
		}
		Info["Domain"] = tname
	}

	b := new(bytes.Buffer)
	for key, value := range Info {
		_, _ = fmt.Fprintf(b, "%-8s: %s\n", key, value)
	}
	result.Result += b.String()

	if len(Nets) > 0 {
		result.Result += fmt.Sprintf("%-8s: %s\n", "Nets", strings.Join(Nets, "\t"))
	}
	result.Result += ntlminfo
	return result
}

type NetbiosInfo struct {
	statusRecv  time.Time
	nameSent    time.Time
	nameRecv    time.Time
	statusReply NetbiosReplyStatus
	nameReply   NetbiosReplyStatus
}

type ProbeNetbios struct {
	socket  net.PacketConn
	replies map[string]*NetbiosInfo
}

type NetbiosReplyHeader struct {
	XID             uint16
	Flags           uint16
	QuestionCount   uint16
	AnswerCount     uint16
	AuthCount       uint16
	AdditionalCount uint16
	QuestionName    [34]byte
	RecordType      uint16
	RecordClass     uint16
	RecordTTL       uint32
	RecordLength    uint16
}

type NetbiosReplyName struct {
	Name [15]byte
	Type uint8
	Flag uint16
}

type NetbiosReplyAddress struct {
	Flag    uint16
	Address [4]uint8
}

var uniqueName []byte

type NetbiosReplyStatus struct {
	Header    NetbiosReplyHeader
	HostName  [15]byte
	UserName  [15]byte
	Names     []NetbiosReplyName
	Addresses []NetbiosReplyAddress
	HWAddr    string
}

func createStatusRequest() []byte {
	return []byte{
		byte(rand.Intn(256)), byte(rand.Intn(256)),
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x20, 0x43, 0x4b, 0x41, 0x41, 0x41,
		0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
		0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
		0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
		0x41, 0x41, 0x41, 0x00, 0x00, 0x21, 0x00, 0x01,
	}
}

func createNameRequest(name string) []byte {
	nbytes := [16]byte{}
	copy(nbytes[0:15], []byte(strings.ToUpper(name)[:]))

	req := []byte{
		byte(rand.Intn(256)), byte(rand.Intn(256)),
		0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x20,
		0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
		0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
		0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
		0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
		0x00, 0x00, 0x20, 0x00, 0x01,
	}

	encoded := encodeNetbiosName(nbytes)
	copy(req[13:45], encoded[0:32])
	return req
}

func encodeNetbiosName(name [16]byte) [32]byte {
	encoded := [32]byte{}

	for i := 0; i < 16; i++ {
		if name[i] == 0 {
			encoded[(i*2)+0] = 'C'
			encoded[(i*2)+1] = 'A'
		} else {
			encoded[(i*2)+0] = byte((name[i] / 16) + 0x41)
			encoded[(i*2)+1] = byte((name[i] % 16) + 0x41)
		}
	}

	return encoded
}

func parseReplay(buff []byte) NetbiosReplyStatus {

	resp := NetbiosReplyStatus{}
	temp := bytes.NewBuffer(buff)

	binary.Read(temp, binary.BigEndian, &resp.Header)
	if resp.Header.QuestionCount != 0 {
		return resp
	}

	if resp.Header.AnswerCount == 0 {
		return resp
	}
	if resp.Header.RecordType == 0x21 {
		var rcnt uint8
		var ridx uint8
		binary.Read(temp, binary.BigEndian, &rcnt)

		for ridx = 0; ridx < rcnt; ridx++ {
			name := NetbiosReplyName{}
			binary.Read(temp, binary.BigEndian, &name)
			resp.Names = append(resp.Names, name)

			if name.Type == 0x20 {
				resp.HostName = name.Name
			}

			if name.Type == 0x03 {
				resp.UserName = name.Name
			}
		}

		var hwbytes [6]uint8
		binary.Read(temp, binary.BigEndian, &hwbytes)
		resp.HWAddr = fmt.Sprintf("%.2x:%.2x:%.2x:%.2x:%.2x:%.2x",
			hwbytes[0], hwbytes[1], hwbytes[2], hwbytes[3], hwbytes[4], hwbytes[5],
		)

		if resp.Header.RecordType == 0x20 {
			var ridx uint16
			for ridx = 0; ridx < (resp.Header.RecordLength / 6); ridx++ {
				addr := NetbiosReplyAddress{}
				binary.Read(temp, binary.BigEndian, &addr)
				resp.Addresses = append(resp.Addresses, addr)
			}
		}

		return resp
	}

	// Addresses
	if resp.Header.RecordType == 0x20 {
		var ridx uint16
		for ridx = 0; ridx < (resp.Header.RecordLength / 6); ridx++ {
			addr := NetbiosReplyAddress{}
			binary.Read(temp, binary.BigEndian, &addr)
			resp.Addresses = append(resp.Addresses, addr)
		}
	}
	return resp
}

func getNbssPayload(name []byte) []byte {
	var payload0 []byte
	n := netbiosEncode(string(name))
	payload0 = append(payload0, []byte("\x81\x00\x00D ")...)
	payload0 = append(payload0, n...)
	payload0 = append(payload0, []byte("\x00 EOENEBFACACACACACACACACACACACACA\x00")...)
	return payload0
}

func probe139(n Netbios, payload []byte) (s string) {
	realhost := fmt.Sprintf("%s:%v", n.Ip, 139)
	conn, err := net.DialTimeout("tcp", realhost, config.TcpConnTimeout)
	if err != nil {
		return
	}
	conn.Write(payload)
	pkg.ReadBytes(conn)
	_, err = conn.Write(NegotiateSMBv1Data1)
	if err != nil {
		return
	}
	_, _ = pkg.ReadBytes(conn)
	_, err = conn.Write(NegotiateSMBv1Data2)
	if err != nil {
		return
	}

	ret, err := pkg.ReadBytes(conn)

	if err != nil || len(ret) < 45 {
		return
	}

	blob_length := uint16(pkg.Bytes2Uint(ret[43:45], '<'))
	blob_count := uint16(pkg.Bytes2Uint(ret[45:47], '<'))

	gss_native := ret[47:]
	off_ntlm := bytes.Index(gss_native, []byte("NTLMSSP"))
	native := gss_native[int(blob_length):blob_count]
	ss := strings.Split(string(native), "\x00\x00")
	//fmt.Println(ss)

	bs := gss_native[off_ntlm:blob_length]
	type2 := ntlmssp.ChallengeMsg{}
	//fmt.Printf("%x\n", bs)
	tinfo := type2.String(bs)
	//fmt.Println(tinfo)

	NativeOS := pkg.TrimName(ss[0])
	NativeLM := pkg.TrimName(ss[1])
	//fmt.Println(NativeOS, NativeLM)
	tinfo += fmt.Sprintf("NativeOS: %s\nNativeLM: %s\n", NativeOS, NativeLM)
	return tinfo
}

func netbiosEncode(name string) (output []byte) {
	var names []int
	src := fmt.Sprintf("%-16s", name)
	for _, a := range src {
		char_ord := int(a)
		high_4_bits := char_ord >> 4
		low_4_bits := char_ord & 0x0f
		names = append(names, high_4_bits, low_4_bits)
	}
	for _, one := range names {
		out := (one + 0x41)
		output = append(output, byte(out))
	}
	return
}

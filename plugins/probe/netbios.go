package probe

import (
	"bytes"
	"cube/model"
	"cube/util"
	"encoding/binary"
	"fmt"
	manuf "github.com/JKme/gomanuf"
	"math/rand"
	"net"
	"strings"
	"time"
)

//From https://github.com/hdm/nextnet

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

type NetbiosReplyStatus struct {
	Header    NetbiosReplyHeader
	HostName  [15]byte
	UserName  [15]byte
	Names     []NetbiosReplyName
	Addresses []NetbiosReplyAddress
	HWAddr    string
}

func NetbiosProbe(task model.ProbeTask) (result model.ProbeTaskResult) {
	result = model.ProbeTaskResult{ProbeTask: task, Result: "", Err: nil}

	conn, err := net.DialTimeout("udp", fmt.Sprintf("%s:%v", task.Ip, task.Port), model.ConnectTimeout)
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		//log.Println("SetReadDeadline failed:", err)
		return
	}
	_, err = conn.Write(createStatusRequest())
	if err != nil {
		return
	}
	ret2, err := util.ReadBytes(conn)
	if err != nil {
		return
	}
	//fmt.Println(text)
	sreply := parseReplay(ret2)
	if len(sreply.Names) == 0 && len(sreply.Addresses) == 0 {
		return
	}

	var nreply NetbiosReplyStatus

	if sreply.Header.RecordType == 0x21 {
		_, err = conn.Write(createNameRequest(util.TrimName(string(sreply.HostName[:]))))
		ret2, _ = util.ReadBytes(conn)
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

	Name, _ := util.ByteToString(sreply.HostName[:])
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

	username := util.TrimName(string(sreply.UserName[:]))
	if len(username) > 0 && username != Info["Name"] {
		Info["Username"] = username
	}

	for _, rname := range sreply.Names {

		//tname := util.TrimName(string(rname.Name[:]))
		tname, _ := util.ByteToString(rname.Name[:])
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
		result.Result += fmt.Sprintf("%-8s: %s", "Nets", strings.Join(Nets, "\t"))
	}
	return result
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

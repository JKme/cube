package probemodule

import (
	"bytes"
	"cube/config"
	"cube/pkg"
	"fmt"
	"github.com/JKme/go-ntlmssp"
	"net"
)

type Wmi struct {
	*Probe
}

func (w Wmi) ProbeName() string {
	return "wmi"
}

func (w Wmi) ProbePort() string {
	return "135"
}

func (w Wmi) PortCheck() bool {
	return true
}

var payload = []byte{5, 0, 11, 3, 16, 0, 0, 0, 120, 0, 40, 0, 3, 0, 0, 0, 184, 16, 184, 16, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 160, 1, 0, 0, 0, 0, 0, 0, 192, 0, 0, 0, 0, 0, 0, 70, 0, 0, 0, 0, 4, 93, 136, 138, 235, 28, 201, 17, 159, 232, 8, 0, 43, 16, 72, 96, 2, 0, 0, 0, 10, 2, 0, 0, 0, 0, 0, 0, 78, 84, 76, 77, 83, 83, 80, 0, 1, 0, 0, 0, 7, 130, 8, 162, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 1, 177, 29, 0, 0, 0, 15}

func (w Wmi) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *w.Probe, Result: "", Err: nil}

	host := fmt.Sprintf("%s:%v", w.Ip, w.Port)
	conn, err := net.DialTimeout("tcp", host, config.TcpConnTimeout)
	if err != nil {
		return result
	}
	conn.Write(payload)
	//if err != nil {
	//	return
	//}
	ret, _ := pkg.ReadBytes(conn)

	off_ntlm := bytes.Index(ret, []byte("NTLMSSP"))
	if off_ntlm == -1 {
		return result
	}
	type2 := ntlmssp.ChallengeMsg{}
	tinfo := type2.String(ret[off_ntlm:])
	result.Result = tinfo
	return result
}

func init() {
	AddProbeKeys("wmi")
}

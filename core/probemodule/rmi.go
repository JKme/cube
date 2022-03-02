package probemodule

import (
	"cube/config"
	"cube/pkg"
	"encoding/hex"
	"fmt"
	"net"
)

type Rmi struct {
	*Probe
}

func (r Rmi) ProbeName() string {
	return "rmi"
}

func (r Rmi) ProbePort() string {
	return "1099"
}

func (r Rmi) ProbeSkipPortCheck() bool {
	return false
}

func (r Rmi) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *r.Probe, Result: "", Err: nil}

	host := fmt.Sprintf("%s:%v", r.Ip, r.Port)
	conn, _ := net.DialTimeout("tcp", host, config.TcpConnTimeout)
	_, err := conn.Write([]byte{0x4a, 0x52, 0x4d, 0x49, 0x00, 0x02, 0x4b})
	if err != nil {
		return result
	}
	r1, _ := pkg.ReadBytes(conn)
	//fmt.Printf("%x", r1[:1])
	if hex.EncodeToString(r1[:1]) == "4e" {
		result.Result = "RMI Registry Deserialization"
	}

	return result
}

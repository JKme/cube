package probemodule

import (
	"cube/config"
	"cube/pkg"
	"fmt"
	"net"
)

type Dubbo struct {
	*Probe
}

func (d Dubbo) ProbeName() string {
	return "dubbo"
}

func (d Dubbo) ProbePort() string {
	return "20880"
}

func (d Dubbo) PortCheck() bool {
	return true
}

func (d Dubbo) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *d.Probe, Result: "", Err: nil}

	host := fmt.Sprintf("%s:%v", d.Ip, d.Port)
	conn, _ := net.DialTimeout("tcp", host, config.TcpConnTimeout)
	_, err := conn.Write([]byte("\r\n\r\n"))
	if err != nil {
		return result
	}
	r1, _ := pkg.ReadBytes(conn)
	//fmt.Printf("Receive: %s\n", string(r1[:5]))
	if string(r1[:5]) == "dubbo" {
		result.Result = "Dubbo Service"
	}

	return result
}
func init() {
	AddProbeKeys("dubbo")
}

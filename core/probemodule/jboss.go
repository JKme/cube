package probemodule

import (
	"cube/config"
	"cube/pkg"
	"encoding/hex"
	"fmt"
	"net"
)

type JBoss struct {
	*Probe
}

func (J JBoss) ProbeName() string {
	return "jboss"
}

func (J JBoss) ProbePort() string {
	return "3873"
}

func (J JBoss) PortCheck() bool {
	return true
}

func (J JBoss) ProbeExec() ProbeResult {
	//https://jspin.re/jboss-eap-as-6-rce-a-little-bit-beyond-xac-xed/
	//https://s3.amazonaws.com/files.joaomatosf.com/slides/alligator_slides.pdf
	result := ProbeResult{Probe: *J.Probe, Result: "", Err: nil}

	host := fmt.Sprintf("%s:%v", J.Ip, J.Port)
	conn, _ := net.DialTimeout("tcp", host, config.TcpConnTimeout)
	//_, err := conn.Write([]byte{0x4a, 0x52, 0x4d, 0x49, 0x00, 0x02, 0x4b})
	//if err != nil {
	//	return result
	//}
	r1, _ := pkg.ReadBytes(conn)
	//fmt.Printf("Receive: %s\n", hex.EncodeToString(r1[:4]))
	if hex.EncodeToString(r1[:4]) == "aced0005" {
		result.Result = "JBoss EAP/AS <= 6.* RCE"
	}
	return result
}

func init() {
	AddProbeKeys("jboss")
}

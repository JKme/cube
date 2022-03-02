package probemodule

import (
	"cube/config"
	"fmt"
	"net"
	"strings"
)

type Zookeeper struct {
	*Probe
}

func (z Zookeeper) ProbeName() string {
	return "zookeeper"
}

func (z Zookeeper) ProbePort() string {
	return "2181"
}

func (z Zookeeper) ProbeSkipPortCheck() bool {
	return false
}

func (z Zookeeper) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *z.Probe, Result: "", Err: nil}

	dl := net.Dialer{Timeout: config.TcpConnTimeout}
	t := fmt.Sprintf("%s:%d", z.Ip, z.Port)
	conn, err := dl.Dial("tcp", t)
	if err != nil {
		result.Err = err
		//log.Printf("Oxid Running Error: %s:%s", task.Ip, err)
		return result
	}
	conn.Write([]byte("version\r\n"))

	r := make([]byte, 1024)
	conn.Read(r)
	if strings.Contains(string(r), "ZooKeeper") {
		result.Result = "Zookeeper Unauthorized"
	}
	return result
}

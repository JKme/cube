package probemodule

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

type Etcd struct {
	*Probe
}

func (e Etcd) ProbeName() string {
	return "etcd"
}

func (e Etcd) ProbePort() string {
	return "2379"
}

func (e Etcd) PortCheck() bool {
	return true
}

func (e Etcd) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *e.Probe, Result: "", Err: nil}

	clt := http.Client{}
	host := fmt.Sprintf("http://%s:%s/version", e.Ip, e.Port)
	req, _ := http.NewRequest("GET", host, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Accept-Charset", "utf-8")
	resp, err := clt.Do(req)
	if err != nil {
		panic(err)
	}
	data := make([]byte, 50)
	c := bufio.NewReader(resp.Body)
	c.Read(data)
	resp.Body.Close()
	if strings.Contains(string(data), "etcd") {
		result.Result = fmt.Sprintf("Etcd Found: %s", string(data))
	}
	return result
}

func init() {
	AddProbeKeys("etcd")
}

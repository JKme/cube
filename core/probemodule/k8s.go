package probemodule

import (
	"bufio"
	"cube/config"
	"fmt"
	"net/http"
	"strings"
)

type K8s struct {
	*Probe
}

func (k K8s) ProbeName() string {
	return "k8s"
}

func (k K8s) ProbePort() string {
	return "10255"
}

func (k K8s) PortCheck() bool {
	return true
}

func (k K8s) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *k.Probe, Result: "", Err: nil}

	clt := http.Client{Timeout: config.TcpConnTimeout}
	host := fmt.Sprintf("http://%s:%s/pods", k.Ip, k.Port)
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
	if strings.Contains(string(data), "PodList") {
		result.Result = fmt.Sprintf("Kubelet Found: %s", string(data))
	}
	return result
}

func init() {
	AddProbeKeys("k8s")
}

package probemodule

import (
	"bufio"
	"crypto/tls"
	"cube/config"
	"fmt"
	"net/http"
	"strings"
)

type K8s10250 struct {
	*Probe
}

func (k K8s10250) ProbeName() string {
	return "k8s10250"
}

func (k K8s10250) ProbePort() string {
	return "10250"
}

func (k K8s10250) PortCheck() bool {
	return true
}

func (k K8s10250) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *k.Probe, Result: "", Err: nil}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	clt := http.Client{Timeout: config.TcpConnTimeout, Transport: tr}
	host := fmt.Sprintf("https://%s:%s/pods", k.Ip, k.Port)
	req, _ := http.NewRequest("GET", host, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Accept-Charset", "utf-8")
	resp, err := clt.Do(req)
	if err != nil {
		panic(err)
	}
	data := make([]byte, 1024)
	c := bufio.NewReader(resp.Body)
	c.Read(data)
	resp.Body.Close()
	if strings.Contains(string(data), "PodList") {
		result.Result = fmt.Sprintf("K8S Vuln Found: K8S Unauthorized Node")
	}
	if resp.StatusCode == 401 && strings.Contains(string(data), "Unauthorized") {
		result.Result = fmt.Sprintf("K8S Node Found, But Need Authorized :(")
	}
	return result
}

func init() {
	AddProbeKeys("k8s10250")
}

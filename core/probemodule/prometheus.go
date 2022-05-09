package probemodule

import (
	"cube/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Prometheus struct {
	*Probe
}

func (p Prometheus) ProbeName() string {
	return "Prometheus"
}

func (p Prometheus) ProbePort() string {
	return "9090"
}

func (p Prometheus) PortCheck() bool {
	return true
}

func (p Prometheus) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *p.Probe, Result: "", Err: nil}

	clt := http.Client{Timeout: config.TcpConnTimeout}
	host := fmt.Sprintf("http://%s:%s/api/v1/status/buildinfo", p.Ip, p.Port)
	req, _ := http.NewRequest("GET", host, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Accept-Charset", "utf-8")
	resp, err := clt.Do(req)
	if err != nil {
		panic(err)
	}
	//data := make([]byte, 1024)
	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	var jsonMap map[string]interface{}
	err = dec.Decode(&jsonMap)

	if err != nil {
		return result
	}
	result.Result = fmt.Sprintf("[*]: %v\n", jsonMap["data"])
	var endpoings = []string{"/api/v1/status/config", "/targets", "/api/v1/status/flags"}
	result.Result += strings.Join(endpoings, "\n")
	return result

}

func init() {
	AddProbeKeys("prometheus")
}

package probemodule

import (
	"bufio"
	"fmt"
	"net/http"
)

type Docker struct {
	*Probe
}

func (d Docker) ProbeName() string {
	return "docker"
}

func (d Docker) ProbePort() string {
	return "2375"
}

func (d Docker) ProbeSkipPortCheck() bool {
	return true
}

func (d Docker) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *d.Probe, Result: "", Err: nil}

	clt := http.Client{}
	host := fmt.Sprintf("http://%s:%s/_ping", d.Ip, d.Port)
	req, _ := http.NewRequest("GET", host, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Accept-Charset", "utf-8")
	resp, err := clt.Do(req)
	if err != nil {
		panic(err)
	}
	data := make([]byte, 2)
	c := bufio.NewReader(resp.Body)
	c.Read(data)
	resp.Body.Close()
	if string(data) == "OK" {
		result.Result = "Docker Remote API Unauthorized Access"
	}
	return result
}

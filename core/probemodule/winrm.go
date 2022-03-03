package probemodule

import (
	"encoding/base64"
	"fmt"
	"github.com/JKme/go-ntlmssp"
	"net/http"
)

type Winrm struct {
	*Probe
}

func (w Winrm) ProbeName() string {
	return "winrm"
}

func (w Winrm) ProbePort() string {
	return "5985"
}

func (w Winrm) PortCheck() bool {
	return true
}

func (w Winrm) ProbeExec() ProbeResult {
	result := ProbeResult{Probe: *w.Probe, Result: "", Err: nil}
	uri := fmt.Sprintf("http://%s:%s/wsman", w.Ip, w.Port)
	clt := http.Client{}
	req, _ := http.NewRequest("POST", uri, nil)
	req.Header.Add("Content-Length", "0")
	req.Header.Add("Keep-Alive", "true")
	req.Header.Add("Content-Type", "application/soap+xml;charset=UTF-8")
	req.Header.Add("User-Agent", "Microsoft WinRM Client")
	req.Header.Add("Authorization", "Negotiate TlRMTVNTUAABAAAAB4IIogAAAAAAAAAAAAAAAAAAAAAGAbEdAAAADw==")
	resp, err := clt.Do(req)
	if err != nil {
		return result
	}
	defer resp.Body.Close()

	ntlminfo := resp.Header.Get("Www-Authenticate")[10:]
	data, err := base64.StdEncoding.DecodeString(ntlminfo)
	type2 := ntlmssp.ChallengeMsg{}
	tinfo := type2.String(data)
	result.Result = tinfo

	return result
}

func init() {
	AddProbeKeys("winrm")
}

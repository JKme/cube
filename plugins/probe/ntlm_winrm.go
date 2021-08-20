package probe

import (
	"cube/model"
	"encoding/base64"
	"fmt"
	"github.com/JKme/go-ntlmssp"
	"net/http"
)

func WinrmProbe(task model.ProbeTask) (result model.ProbeTaskResult) {
	result = model.ProbeTaskResult{ProbeTask: task, Result: "", Err: nil}
	uri := fmt.Sprintf("http://%s:%s/wsman", task.Ip, task.Port)
	clt := http.Client{}
	req, _ := http.NewRequest("POST", uri, nil)
	req.Header.Add("Content-Length", "0")
	req.Header.Add("Keep-Alive", "true")
	req.Header.Add("Content-Type", "application/soap+xml;charset=UTF-8")
	req.Header.Add("User-Agent", "Microsoft WinRM Client")
	req.Header.Add("Authorization", "Negotiate TlRMTVNTUAABAAAAB4IIogAAAAAAAAAAAAAAAAAAAAAGAbEdAAAADw==")
	resp, err := clt.Do(req)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	ntlminfo := resp.Header.Get("Www-Authenticate")[10:]
	data, err := base64.StdEncoding.DecodeString(ntlminfo)
	type2 := ntlmssp.ChallengeMsg{}
	tinfo := "\n" + type2.String(data)
	result.Result = tinfo

	return result
}

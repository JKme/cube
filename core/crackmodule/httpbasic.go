package crackmodule

import (
	"cube/config"
	"cube/gologger"
	"net/http"
	"strings"
)

type HttpBasic struct {
	*Crack
}

func (h HttpBasic) CrackName() string {
	return "HttpBasic"
}

func (h HttpBasic) CrackPort() string {
	return "80"
}

func (h HttpBasic) CrackAuthUser() []string {
	return []string{"root", "admin", "tomcat", "test", "guest"}
}

func (h HttpBasic) CrackAuthPass() []string {
	return config.PASSWORDS
}

func (h HttpBasic) IsMutex() bool {
	return true
}

func (h HttpBasic) CrackPortCheck() bool {
	return false
}

func (h HttpBasic) Exec() CrackResult {
	result := CrackResult{Crack: *h.Crack, Result: false, Err: nil}

	clt := http.Client{}
	if !strings.HasPrefix(h.Ip, "http") {
		gologger.Errorf("Invalid URL, eg: http://%s", h.Ip)
	}
	req, _ := http.NewRequest("POST", h.Ip, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.SetBasicAuth(h.Auth.User, h.Auth.Password)
	res, err := clt.Do(req)
	if err != nil {
		gologger.Error(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 401 {
		result.Result = true
	}
	return result
}

func init() {
	AddCrackKeys("httpbasic")
}

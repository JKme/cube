package crackmodule

import (
	"crypto/tls"
	"cube/config"
	"cube/gologger"
	"log"
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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	clt := http.Client{Transport: tr}
	if !strings.HasPrefix(h.Ip, "http") {
		gologger.Errorf("Invalid URL, eg: http://%s", h.Ip)
	}
	req, _ := http.NewRequest("POST", h.Ip, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.SetBasicAuth(h.Auth.User, h.Auth.Password)
	res, err := clt.Do(req)

	if err != nil {
		// 处理请求错误
		log.Printf("Error making request: %v", err)
		// 应该在这里返回或处理错误
		return result
	}
	if res != nil {
		defer func() {
			// 使用 defer 调用匿名函数来处理 Close 的错误
			if err := res.Body.Close(); err != nil {
				// 处理关闭 resp.Body 时的错误
				log.Printf("Error closing response body: %v", err)
			}
		}()
		if res.StatusCode >= 200 && res.StatusCode < 400 {
			result.Result = true
		}
	} else {
		// 如果到这里，说明有严重的错误发生，resp2 应该不为 nil。
		log.Printf("Response is nil without a preceding error.")
	}

	//if res.StatusCode != 401 {
	//	result.Result = true
	//}
	return result
}

func init() {
	AddCrackKeys("httpbasic")
}

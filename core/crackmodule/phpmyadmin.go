package crackmodule

import (
	"bufio"
	"cube/config"
	"cube/gologger"
	"fmt"
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

func (p Phpmyadmin) SkipPortCheck() bool {
	return true
}

func (p Phpmyadmin) Exec() CrackResult {
	result := CrackResult{Crack: *p.Crack, Result: "", Err: nil}

	clt := http.Client{}
	if !strings.HasPrefix(p.Ip, "http://") {
		gologger.Errorf("Invalid URL, eg: http://%s", p.Ip)
	}
	req, _ := http.NewRequest("GET", p.Ip, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Accept-Charset", "utf-8")
	resp, err := clt.Do(req)
	if err != nil {
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
		Jar: jar}

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

	resp2, _ := crackClt.Do(req2)
	//resp2, _ := crackClt.PostForm(task.Ip, urlValues)
	//resp2, _ := crackClt.Post(task.Ip, urlValues)
	defer resp2.Body.Close()
	if resp2.StatusCode == 302 {
		result.Result = fmt.Sprintf("User: %s \tPassword: %s", p.Auth.User, p.Auth.Password)
	}

	return result
}

func init() {
	AddCrackKeys("phpmyadmin")
}

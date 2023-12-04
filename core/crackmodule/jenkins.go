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

type Jenkins struct {
	*Crack
}

func (j Jenkins) CrackName() string {
	return "jenkins"
}

func (j Jenkins) CrackPort() string {
	return "80"
}

func (j Jenkins) CrackAuthUser() []string {
	return []string{"jenkins", "admin"}
}

func (j Jenkins) CrackAuthPass() []string {
	return config.PASSWORDS
}

func (j Jenkins) IsMutex() bool {
	return true
}

func (j Jenkins) CrackPortCheck() bool {
	return true
}

func (j Jenkins) Exec() CrackResult {
	result := CrackResult{Crack: *j.Crack, Result: false, Err: nil}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	clt := http.Client{Transport: tr}
	if !strings.HasPrefix(j.Ip, "http") {
		gologger.Errorf("Invalid URL, eg: http://%s", j.Ip)
	}
	req, _ := http.NewRequest("GET", j.Ip+"/login", nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	resp, err := clt.Do(req)
	if err != nil {
		// 处理请求错误
		log.Printf("Error making request: %v", err)
		// 应该在这里返回或处理错误
		return result
	}

	data := make([]byte, 20250)
	c := bufio.NewReader(resp.Body)
	c.Read(data)
	resp.Body.Close()
	//content, _ := ioutil.ReadAll(resp.Body)

	r := regexp.MustCompile(`(?U)action="(.*)"`)
	match := r.FindStringSubmatch(string(data))
	if match == nil {
		return result
	}
	postUri := strings.TrimSpace(match[1])
	//fmt.Println(postUri)

	//clt2 := http.Client{
	//	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	//		return http.ErrUseLastResponse
	//	},
	//}

	jar, _ := cookiejar.New(nil)
	host, _ := url.Parse(j.Ip)
	jar.SetCookies(host, resp.Cookies())

	clt2 := http.Client{
		Jar:       jar,
		Transport: tr,
	}
	urlValues := url.Values{}
	urlValues.Add("j_username", j.Auth.User)
	urlValues.Add("j_password", j.Auth.Password)
	body := strings.NewReader(urlValues.Encode())
	req2, _ := http.NewRequest("POST", j.Ip+"/"+postUri, body)
	req2.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req2.Header.Add("Connection", "close")
	req2.Header.Add("Accept-Charset", "utf-8")
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r2, err := clt2.Do(req2)
	if err != nil {
		// 处理请求错误
		log.Printf("Error making request: %v", err)
		// 应该在这里返回或处理错误
		return result
	}

	defer r2.Body.Close()
	data2 := make([]byte, 10480)
	c2 := bufio.NewReader(r2.Body)
	c2.Read(data2)
	//fmt.Println(string(data2))
	//fmt.Print(r2.Header["Set-Cookie"])
	if strings.Contains(string(data2), "Dashboard") {
		result.Result = true
	}
	return result
}

func init() {
	AddCrackKeys("jenkins")
}

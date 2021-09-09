package crack

import (
	"bufio"
	"cube/model"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

func ZabbixCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	clt := http.Client{}
	req, _ := http.NewRequest("GET", task.Ip, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Accept-Charset", "utf-8")
	resp, err := clt.Do(req)
	if err != nil {
		panic(err)
	}

	data := make([]byte, 20250)
	c := bufio.NewReader(resp.Body)
	c.Read(data)
	resp.Body.Close()
	r := regexp.MustCompile(`(?U)name="sid" value="(.*)"`)
	match := r.FindStringSubmatch(string(data))
	if match == nil {
		return
	}
	token := strings.TrimSpace(match[1])

	jar, _ := cookiejar.New(nil)
	host, _ := url.Parse(task.Ip)
	jar.SetCookies(host, resp.Cookies())
	crackClt := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar: jar}

	//fmt.Println(jar.Cookies(host))

	urlValues := url.Values{}
	urlValues.Add("name", task.Auth.User)
	urlValues.Add("password", task.Auth.Password)
	urlValues.Add("form_refresh", "1")
	urlValues.Add("enter", "Sign in")
	urlValues.Add("sid", token)

	body := strings.NewReader(urlValues.Encode())
	req2, _ := http.NewRequest("POST", task.Ip, body)
	req2.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req2.Header.Add("Connection", "close")
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp2, _ := crackClt.Do(req2)
	defer resp2.Body.Close()
	if resp2.StatusCode == 302 {
		result.Result = fmt.Sprintf("User: %s \tPassword: %s", task.Auth.User, task.Auth.Password)
	}

	return result
}

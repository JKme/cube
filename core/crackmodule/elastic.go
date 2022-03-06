package crackmodule

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Elastic struct {
	*Crack
}

func (e Elastic) CrackName() string {
	return "elastic"
}

func (e Elastic) CrackPort() string {
	return "9200"
}

func (e Elastic) CrackAuthUser() []string {
	return []string{""}
}

func (e Elastic) CrackAuthPass() []string {
	return []string{""}
}

func (e Elastic) IsMutex() bool {
	return false
}

func (e Elastic) CrackPortCheck() bool {
	return true
}

func (e Elastic) Exec() CrackResult {
	result := CrackResult{Crack: *e.Crack, Result: "", Err: nil}

	url := fmt.Sprintf("http://%s:%v/_cat", e.Ip, e.Port)
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result
	}
	res.Header.Add("User-agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	res.Header.Add("Accept", "*/*")
	res.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	res.Header.Add("Connection", "close")

	clt := http.Client{}
	resp, err := clt.Do(res)
	if err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if strings.Contains(string(body), "/_cat/master") {
			//result.Result = fmt.Sprintf("User: %s\tPassword: %s \t", task.Auth.User, task.Auth.Password)
			result.Result = fmt.Sprintf("Elasticsearch unauthorized")

		}
	}
	return result
}

func init() {
	AddCrackKeys("elastic")
}

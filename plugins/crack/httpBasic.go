package crack

import (
	"cube/model"
	"fmt"
	"net/http"
)

func HttpBasicCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	clt := http.Client{}
	req, _ := http.NewRequest("POST", task.Ip, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Add("Connection", "close")
	req.SetBasicAuth(task.Auth.User, task.Auth.Password)
	res, err := clt.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 401 {
		result.Result = fmt.Sprintf("Password: %s", task.Auth.Password)
	}
	return result
}

package crack

import (
	"cube/model"
	"fmt"
	"github.com/jlaffaye/ftp"
)

func FtpCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	conn, err := ftp.DialTimeout(fmt.Sprintf("%v:%v", task.Ip, task.Port), model.ConnectTimeout)
	if err == nil {
		err = conn.Login(task.Auth.User, task.Auth.Password)
		if err == nil {
			defer conn.Logout()
			result.Result = fmt.Sprintf("User: %s\tPassword: %s \t", task.Auth.User, task.Auth.Password)
		}
	}
	return result
}

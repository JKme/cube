package crack

import (
	"cube/model"
	"cube/util"
	"fmt"
	"github.com/stacktitan/smb/smb"
	"strconv"
)

func SmbCrack(task model.CrackTask) (result model.CrackTaskResult) {
	Port, _ := strconv.Atoi(task.Port)
	options := smb.Options{
		Host:        task.Ip,
		Port:        Port,
		User:        task.Auth.User,
		Password:    task.Auth.Password,
		Domain:      "",
		Workstation: "",
	}
	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			result.Result = util.Green(fmt.Sprintf("User: %s\tPassword: %s \t", task.Auth.User, task.Auth.Password))
		}
	}
	return result
}

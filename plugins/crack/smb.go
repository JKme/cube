package crack

import (
	"cube/model"
	"fmt"
	"github.com/stacktitan/smb/smb"
	"strconv"
	"strings"
)

func SmbCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	Port, _ := strconv.Atoi(task.Port)
	User := task.Auth.User
	Domain := ""
	if strings.Contains(User, "\\") {
		l := strings.Split(User, "\\")
		Domain, User = l[0], l[1]
	}
	options := smb.Options{
		Host:        task.Ip,
		Port:        Port,
		User:        User,
		Password:    task.Auth.Password,
		Domain:      Domain,
		Workstation: "",
	}
	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			result.Result = fmt.Sprintf("User: %s\tPassword: %s \t", task.Auth.User, task.Auth.Password)
		}
	}
	return result
}

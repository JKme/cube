package crack

import (
	"cube/model"
	"fmt"
	"golang.org/x/crypto/ssh"
)

func SshCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	config := &ssh.ClientConfig{
		User: task.Auth.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(task.Auth.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", task.Ip, task.Port), config)
	if err == nil {
		defer conn.Close()
		session, err := conn.NewSession()
		errRet := session.Run("echo Hello")
		if err == nil && errRet == nil {
			defer session.Close()
			result.Result = fmt.Sprintf("User: %s \t Password: %s", task.Auth.User, task.Auth.Password)

		}
	}
	return result
}

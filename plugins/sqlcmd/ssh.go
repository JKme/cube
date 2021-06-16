package sqlcmd

import (
	"cube/log"
	"cube/model"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
)

func SshCmd(task model.SqlcmdTask) (result model.SqlcmdTaskResult) {
	result = model.SqlcmdTaskResult{SqlcmdTask: task, Result: "", Err: nil}
	config := &ssh.ClientConfig{
		//Timeout: time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User: task.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(task.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", task.Ip, task.Port), config)

	if err == nil {
		defer conn.Close()
		session, err := conn.NewSession()
		r, err := session.Output(task.Query)
		result.Result = string(r)
		if err != nil {
			log.Fatalf("Failed to run command, Err:%s", err.Error())
			os.Exit(0)
		}
	}
	return result
}

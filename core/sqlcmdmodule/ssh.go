package sqlcmdmodule

import (
	"cube/gologger"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
)

type CmdSsh struct {
	*Sqlcmd
}

func (c CmdSsh) SqlcmdName() string {
	return "ssh"
}

func (c CmdSsh) SqlcmdPort() string {
	return "22"
}

func (c CmdSsh) SqlcmdExec() SqlcmdResult {
	result := SqlcmdResult{Sqlcmd: *c.Sqlcmd, Result: "", Err: nil}

	config := &ssh.ClientConfig{
		//Timeout: time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User: c.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", c.Ip, c.Port), config)

	if err == nil {
		defer conn.Close()
		session, err := conn.NewSession()
		r, err := session.Output(c.Query)
		result.Result = string(r)
		if err != nil {
			gologger.Error("Failed to run command, Err:%s", err.Error())
			os.Exit(0)
		}
	}
	return result
}

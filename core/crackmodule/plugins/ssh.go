package plugins

import (
	"cube/conf"
	"fmt"
	"golang.org/x/crypto/ssh"
)

type SshCrack struct {
	*Crack
}

func (sshCrack SshCrack) SetName() string {
	return "ssh"
}

func (sshCrack SshCrack) SetPort() string {
	return "22"
}

func (sshCrack SshCrack) SetAuthUser() []string {
	return []string{"root", "admin"}
}

func (sshCrack SshCrack) SetAuthPass() []string {
	return conf.PASSWORDS
}

func (sshCrack SshCrack) IsLoad() bool {
	return true
}

func (sshCrack SshCrack) IsTcp() bool {
	return true
}

func (sshCrack SshCrack) IsMutex() bool {
	return false
}

func (sshCrack SshCrack) Exec() (crackResult CrackResult) {
	crackResult = CrackResult{Crack: *sshCrack.Crack, Result: "", Err: nil}
	config := &ssh.ClientConfig{
		User: sshCrack.Auth.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshCrack.Auth.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", sshCrack.Ip, sshCrack.Port), config)
	if err != nil {
		crackResult.Err = err
		return
	}
	defer conn.Close()
	session, err := conn.NewSession()
	errRet := session.Run("echo Hello")
	if err == nil && errRet == nil {
		defer session.Close()
		crackResult.Result = fmt.Sprintf("User: %s \t Password: %s", sshCrack.Auth.User, sshCrack.Auth.Password)

	}
	return crackResult
}

func init() {
	AddCrackKeys("ssh")
}

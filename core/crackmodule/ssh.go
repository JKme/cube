package crackmodule

import (
	"cube/config"
	"fmt"
	"golang.org/x/crypto/ssh"
)

type SshCrack struct {
	*Crack
}

func (sshCrack SshCrack) CrackName() string {
	return "ssh"
}

func (sshCrack SshCrack) CrackPort() string {
	return "22"
}

func (sshCrack SshCrack) CrackAuthUser() []string {
	return []string{"root", "admin"}
}

func (sshCrack SshCrack) CrackAuthPass() []string {
	return config.PASSWORDS
}

func (sshCrack SshCrack) CrackPortCheck() bool {
	return true
}

func (sshCrack SshCrack) IsMutex() bool {
	return false
}

func (sshCrack SshCrack) Exec() (crackResult CrackResult) {
	crackResult = CrackResult{Crack: *sshCrack.Crack, Result: false, Err: nil}
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
		crackResult.Result = true

	}
	return crackResult
}

func init() {
	AddCrackKeys("ssh")
}

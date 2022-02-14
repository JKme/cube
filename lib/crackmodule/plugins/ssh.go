package plugins

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

type SshCrack struct {
	*Crack
}

//type SshCrack = Crack

func (sshCrack SshCrack) SetName() (s string) {
	return "ssh"
}

func (sshCrack SshCrack) Desc() (s string) {
	return "crack ssh service password"
}
func (sshCrack SshCrack) Load() (b bool) {
	return true
}
func (sshCrack SshCrack) GetPort() (s string) {
	return "22"
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

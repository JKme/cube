package crackmodule

import (
	"cube/config"
	"fmt"
	"github.com/jlaffaye/ftp"
)

type FtpCrack struct {
	*Crack
}

func (f FtpCrack) CrackName() string {
	return "ftp"
}

func (f FtpCrack) CrackPort() string {
	return "21"
}

func (f FtpCrack) CrackAuthUser() []string {
	return []string{"anonymous", "ftp", "admin", "www", "web", "root", "db", "wwwroot", "data"}
}

func (f FtpCrack) CrackAuthPass() []string {
	return config.PASSWORDS
}

func (f FtpCrack) IsMutex() bool {
	return false
}

func (f FtpCrack) CrackPortCheck() bool {
	return true
}

func (f FtpCrack) Exec() (result CrackResult) {
	result = CrackResult{Crack: *f.Crack, Result: "", Err: nil}

	conn, err := ftp.DialTimeout(fmt.Sprintf("%v:%v", f.Ip, f.Port), config.TcpConnTimeout)
	if err == nil {
		err = conn.Login(f.Auth.User, f.Auth.Password)
		if err == nil {
			defer conn.Logout()
			result.Result = fmt.Sprintf("User: %s\tPassword: %s \t", f.Auth.User, f.Auth.Password)
		}
	}
	return result
}

func init() {
	AddCrackKeys("ftp")
}

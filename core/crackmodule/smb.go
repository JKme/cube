package crackmodule

import (
	"cube/config"
	"fmt"
	"github.com/stacktitan/smb/smb"
	"strconv"
	"strings"
)

type Smb struct {
	*Crack
}

func (s Smb) CrackName() string {
	return "smb"
}

func (s Smb) CrackPort() string {
	return "445"
}

func (s Smb) CrackAuthUser() []string {
	return []string{"administrator", "admin", "guest"}
}

func (s Smb) CrackAuthPass() []string {
	return config.PASSWORDS
}

func (s Smb) IsLoad() bool {
	return true
}

func (s Smb) IsMutex() bool {
	return false
}

func (s Smb) IsTcp() bool {
	return true
}

func (s Smb) Exec() CrackResult {
	result := CrackResult{Crack: *s.Crack, Result: "", Err: nil}

	Port, _ := strconv.Atoi(s.Port)
	User := s.Auth.User
	Domain := ""
	if strings.Contains(User, "\\") {
		l := strings.Split(User, "\\")
		Domain, User = l[0], l[1]
	}
	options := smb.Options{
		Host:        s.Ip,
		Port:        Port,
		User:        User,
		Password:    s.Auth.Password,
		Domain:      Domain,
		Workstation: "",
	}
	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			result.Result = fmt.Sprintf("User: %s\tPassword: %s \t", s.Auth.User, s.Auth.Password)
		}
	}
	return result
}

func init() {
	AddCrackKeys("smb")
}

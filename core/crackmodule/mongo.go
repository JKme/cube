package crackmodule

import (
	"cube/config"
	"fmt"
	"gopkg.in/mgo.v2"
)

type Mongo struct {
	*Crack
}

func (m Mongo) CrackName() string {
	return "mongo"
}

func (m Mongo) CrackPort() string {
	return "27017"
}

func (m Mongo) CrackAuthUser() []string {
	return []string{"root", "admin"}
}

func (m Mongo) CrackAuthPass() []string {
	return config.PASSWORDS
}

func (m Mongo) IsMutex() bool {
	return true
}

func (m Mongo) CrackPortCheck() bool {
	return true
}

func (m Mongo) Exec() CrackResult {
	result := CrackResult{Crack: *m.Crack, Result: "", Err: nil}
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", m.Auth.User, m.Auth.Password, m.Ip, m.Port, "test")
	session, err := mgo.DialWithTimeout(url, config.TcpConnTimeout)

	if err == nil {
		defer session.Close()
		err = session.Ping()
		if err == nil {
			result.Result = fmt.Sprintf("User: %s\tPassword: %s \t", m.Auth.User, m.Auth.Password)
		}
	}
	return result
}

func init() {
	AddCrackKeys("mongo")
}

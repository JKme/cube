package crackmodule

import (
	"cube/config"
	"database/sql"
	"fmt"
	_ "github.com/sijms/go-ora/v2"
	"time"
)

// Oracle https://github.com/shadow1ng/fscan/blob/main/Plugins/oracle.go
type Oracle struct {
	*Crack
}

func (o Oracle) CrackName() string {
	return "oracle"
}

func (o Oracle) CrackPort() string {
	return "1521"
}

func (o Oracle) CrackAuthUser() []string {
	return []string{"sys", "system", "admin", "test", "web", "orcl"}
}

func (o Oracle) CrackAuthPass() []string {
	return config.PASSWORDS
}

func (o Oracle) IsMutex() bool {
	return false
}

func (o Oracle) CrackPortCheck() bool {
	return true
}

func (o Oracle) Exec() CrackResult {
	result := CrackResult{Crack: *o.Crack, Result: false, Err: nil}

	Host, Port, Username, Password := o.Ip, o.Port, o.Auth.User, o.Auth.Password
	dataSourceName := fmt.Sprintf("oracle://%s:%s@%s:%s/orcl", Username, Password, Host, Port)
	db, err := sql.Open("oracle", dataSourceName)
	if err == nil {
		db.SetConnMaxLifetime(time.Duration(config.TcpConnTimeout) * time.Second)
		db.SetConnMaxIdleTime(time.Duration(config.TcpConnTimeout) * time.Second)
		db.SetMaxIdleConns(0)
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result.Result = true
		}
	}
	return result
}

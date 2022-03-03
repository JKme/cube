package crackmodule

import (
	"cube/config"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

type Mssql struct {
	*Crack
}

func (m Mssql) CrackName() string {
	return "mssql"
}

func (m Mssql) CrackPort() string {
	return "1433"
}

func (m Mssql) CrackAuthUser() []string {
	return []string{"sa", "sql"}
}

func (m Mssql) CrackAuthPass() []string {
	return config.PASSWORDS
}

func (m Mssql) IsMutex() bool {
	return false
}

func (m Mssql) CrackPortCheck() bool {
	return true
}

func (m Mssql) Exec() CrackResult {
	result := CrackResult{Crack: *m.Crack, Result: "", Err: nil}

	dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", m.Ip,
		m.Port, m.Auth.User, m.Auth.Password, "tempdb")
	db, err := sql.Open("mssql", dataSourceName)
	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result.Result = fmt.Sprintf("User: %s\tPassword: %s \t", m.Auth.User, m.Auth.Password)
		}
	}
	return result
}

func init() {
	AddCrackKeys("mssql")
}

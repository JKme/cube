package crackmodule

import (
	"cube/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Postgres struct {
	*Crack
}

func (p Postgres) CrackName() string {
	return "postgres"
}

func (p Postgres) CrackPort() string {
	return "5432"
}

func (p Postgres) CrackAuthUser() []string {
	return []string{"postgres", "admin", "root"}
}

func (p Postgres) CrackAuthPass() []string {
	return config.PASSWORDS
}

func (p Postgres) IsMutex() bool {
	return false
}

func (p Postgres) CrackPortCheck() bool {
	return true
}

func (p Postgres) Exec() CrackResult {
	result := CrackResult{Crack: *p.Crack, Result: false, Err: nil}

	dataSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", p.Auth.User,
		p.Auth.Password, p.Ip, p.Port, "postgres", "disable")
	db, err := sql.Open("postgres", dataSourceName)

	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result.Result = true
		}
	}
	return result
}

func init() {
	AddCrackKeys("postgres")
}

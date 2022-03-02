package crackmodule

import (
	"cube/config"
	"database/sql"
	"fmt"
	"strings"
)

type Mysql struct {
	*Crack
}

func (m Mysql) CrackName() string {
	return "mysql"
}

func (m Mysql) CrackPort() string {
	return "3306"
}

func (m Mysql) CrackAuthUser() []string {
	return []string{"root", "mysql"}
}

func (m Mysql) CrackAuthPass() []string {
	return config.PASSWORDS
}

func (m Mysql) IsMutex() bool {
	return false
}

func (m Mysql) SkipPortCheck() bool {
	return true
}

func (m Mysql) Exec() CrackResult {
	result := CrackResult{Crack: *m.Crack, Result: "", Err: nil}

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/mysql?charset=utf8&timeout=%v", m.Auth.User, m.Auth.Password, m.Ip, m.Port, config.TcpConnTimeout)
	db, err := sql.Open("mysql", dataSourceName)
	if err == nil {
		err = db.Ping()
		if err == nil {
			rows, err := db.Query("select @@version, @@version_compile_os, @@version_compile_machine, @@secure_file_priv;")
			if err == nil {
				var s string
				cols, _ := rows.Columns()
				for rows.Next() {
					err := rows.Scan(&cols[0], &cols[1], &cols[2], &cols[3])
					if err != nil {
						fmt.Println(err)
					}
					result.Extra = fmt.Sprintf("OS=%s Version=%s Arch=%s File_Priv=%s\t", strings.Split(cols[1], "-")[0], cols[0], cols[2], cols[3])

				}
				result.Result = fmt.Sprintf("User: %s \tPassword: %s \t %s", m.Auth.User, m.Auth.Password, s)
			}
		}
		db.Close()
	}
	return result
}

func init() {
	AddCrackKeys("mysql")
}

package crack

import (
	"cube/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func MysqlCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/mysql?charset=utf8&timeout=%v", task.Auth.User, task.Auth.Password, task.Ip, task.Port, model.ConnectTimeout)
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
					s = s + fmt.Sprintf("OS=%s Version=%s Machine=%s File_Priv=%s\t", strings.Split(cols[1], "-")[0], cols[0], cols[2], cols[3])

				}
				result.Result = fmt.Sprintf("User: %s \tPassword: %s \t %s", task.Auth.User, task.Auth.Password, s)
			}
		}
		db.Close()
	}

	return result
}

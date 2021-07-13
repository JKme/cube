package crack

import (
	"cube/model"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func MssqlCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}

	dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", task.Ip,
		task.Port, task.Auth.User, task.Auth.Password, "tempdb")
	db, err := sql.Open("mssql", dataSourceName)
	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result.Result = fmt.Sprintf("User: %s\tPassword: %s \t", task.Auth.User, task.Auth.Password)
		}
	}

	return result
}

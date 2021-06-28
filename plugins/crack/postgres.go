package crack

import (
	"cube/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func PostgresCrack(task model.CrackTask) (result model.CrackTaskResult) {
	dataSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", task.Auth.User,
		task.Auth.Password, task.Ip, task.Port, "postgres", "disable")
	db, err := sql.Open("postgres", dataSourceName)

	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result.Result = fmt.Sprintf("User: %s\tPassword: %s \t", task.Auth.User, task.Auth.Password)
		}
	}
	return result
}

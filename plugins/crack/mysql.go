package crack

import (
	"cube/model"
	"cube/util"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func MysqlCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	if !util.PortCheck(task.Ip, task.Port) {
		return
	} else {
		result.Result = fmt.Sprintf("Open:%s", task.Port)
	}

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/mysql?charset=utf8&timeout=%v", task.Auth.User, task.Auth.Password, task.Ip, task.Port, time.Duration(3)*time.Second)
	db, err := sql.Open("mysql", dataSourceName)
	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result.Result = util.Green(fmt.Sprintf("User: %s \tPassword: %s", task.Auth.User, task.Auth.Password))
		}
	}
	return result
}

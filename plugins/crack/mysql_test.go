package crack

import (
	"cube/model"
	"fmt"
	"testing"
)

func TestMysqlCrack(t *testing.T) {
	task := model.CrackTask{
		Ip:   "127.0.0.1",
		Port: "3306",
		Auth: model.Auth{
			User:     "root",
			Password: "root",
		},
		CrackPlugin: "mysql",
	}
	r := MysqlCrack(task)
	fmt.Println(r)
}

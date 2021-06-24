package crack

import (
	"cube/model"
	"fmt"
	"testing"
)

func TestMysqlCrack(t *testing.T) {
	task := model.CrackTask{
		Ip:   "172.20.41.254",
		Port: "3306",
		Auth: model.Auth{
			User:     "root",
			Password: "123456",
		},
		CrackPlugin: "mysql",
	}
	r := MysqlCrack(task)
	fmt.Println(r)
}

func TestRedisCrack(t *testing.T) {
	task := model.CrackTask{
		Ip:   "127.0.0.1",
		Port: "6379",
		Auth: model.Auth{
			User:     "",
			Password: "123456",
		},
		CrackPlugin: "redis",
	}
	r := RedisCrack(task)
	fmt.Println(r)
}

func TestSlice(t *testing.T) {
	var UserDict = map[string][]string{
		"ftp":        {"ftp", "admin", "www", "web", "root", "db", "wwwroot", "data"},
		"mysql":      {"root", "mysql"},
		"mssql":      {"sa", "sql"},
		"smb":        {"administrator", "admin", "guest"},
		"postgresql": {"postgres", "admin"},
		"ssh":        {"root", "admin"},
		"mongodb":    {"root", "admin"},
	}
	fmt.Println(UserDict["aa"])
}

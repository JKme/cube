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

func TestFtpCrack(t *testing.T) {
	task := model.CrackTask{
		Ip:   "127.0.0.1",
		Port: "21",
		Auth: model.Auth{
			User:     "root",
			Password: "root",
		},
		CrackPlugin: "ftp",
	}
	r := FtpCrack(task)
	fmt.Println(r)
}

func TestSmbCrack(t *testing.T) {
	task := model.CrackTask{
		Ip:   "172.20.7.20",
		Port: "445",
		Auth: model.Auth{
			User:     "Guest",
			Password: "",
		},
		CrackPlugin: "smb",
	}
	r := SmbCrack(task)
	fmt.Println(r)
}

func TestElasticCrack(t *testing.T) {
	task := model.CrackTask{
		Ip:   "192.168.249.240",
		Port: "9200",
		Auth: model.Auth{
			User:     "",
			Password: "",
		},
		CrackPlugin: "elastic",
	}
	r := ElasticCrack(task)
	fmt.Println(r)
}

func TestPhpmyadminCrack(t *testing.T) {
	task := model.CrackTask{
		Ip:   "http://127.0.0.1:8081/",
		Port: "",
		Auth: model.Auth{
			User:     "root",
			Password: "root1",
		},
		CrackPlugin: "phpmyadmin",
	}
	r := PhpmyadminCrack(task)
	fmt.Println(r)
}

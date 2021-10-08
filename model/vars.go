package model

import (
	"sync"
	"time"
)

const (
	ConnectTimeout = 5 * time.Second
	ThreadTimeout  = 7 * time.Second
)

var (
	CommonPortMap map[string]int
	//SuccessHash   map[string]bool
	Mutex sync.Mutex
)

var SuccessHash = struct {
	sync.RWMutex
	S map[string]bool
}{S: make(map[string]bool)}

var UserDict = map[string][]string{
	"ftp":        {"anonymous", "ftp", "admin", "www", "web", "root", "db", "wwwroot", "data"},
	"mysql":      {"root", "mysql"},
	"mssql":      {"sa", "sql"},
	"smb":        {"administrator", "admin", "guest"},
	"postgres":   {"postgres", "admin"},
	"ssh":        {"root", "admin"},
	"mongodb":    {"root", "admin"},
	"phpmyadmin": {"root"},
	"httpbasic":  {"root", "admin", "tomcat", "test", "guest"}, //activemq、tomcat、nexus
	"elastic":    {""},
	"jenkins":    {"jenkins", "admin"},
	"zabbix":     {"admin", "guest"},
}

var PassDict = []string{" ", "123456", "admin", "admin123", "root", "5201314", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "{user}", "{user}1", "{user}12", "{user}111", "{user}123", "{user}1234", "{user}12345", "{user}123456", "{user}@123", "{user}_123", "{user}#123", "{user}@111", "{user}@2019", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "000000", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "1QAZ2wsx", "1q2w3e4r", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "123456a", "123456aa", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system"}

func init() {
	CommonPortMap = make(map[string]int)
	CommonPortMap["ftp"] = 21
	CommonPortMap["ssh"] = 22
	CommonPortMap["oxid"] = 135
	CommonPortMap["netbios"] = 137
	CommonPortMap["ntlm-smb"] = 445
	CommonPortMap["ntlm-winrm"] = 5985
	CommonPortMap["ntlm-wmi"] = 135
	CommonPortMap["ntlm-mssql"] = 1433
	CommonPortMap["rmi"] = 1099
	CommonPortMap["docker"] = 2375
	CommonPortMap["dubbo"] = 20880

	CommonPortMap["ms17010"] = 445
	CommonPortMap["smbghost"] = 445
	CommonPortMap["mssql"] = 1433
	CommonPortMap["mssql-wscript"] = 1433
	CommonPortMap["mssql-com"] = 1433
	CommonPortMap["mssql-clr"] = 1433
	CommonPortMap["zookeeper"] = 2181
	CommonPortMap["mysql"] = 3306
	CommonPortMap["postgres"] = 5432
	CommonPortMap["redis"] = 6379
	CommonPortMap["elastic"] = 9200
	CommonPortMap["mongo"] = 27017

}

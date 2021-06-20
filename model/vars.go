package model

import (
	"sync"
	"time"
)

const (
	TIMEUNIT = 5
	T        = 6 * time.Second
	TIMEOUT  = time.Duration(TIMEUNIT) * time.Second
)

var (
	CommonPortMap map[string]int
	SuccessHash   map[string]bool
	Mutex         sync.Mutex
)

var UserDict = map[string][]string{
	"ftp":        {"ftp", "admin", "www", "web", "root", "db", "wwwroot", "data"},
	"mysql":      {"root", "mysql"},
	"mssql":      {"sa", "sql"},
	"smb":        {"administrator", "admin", "guest"},
	"postgresql": {"postgres", "admin"},
	"ssh":        {"root", "admin"},
	"mongodb":    {"root", "admin"},
}

var PassDict = []string{"123456", "admin", "admin123", "root", "", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "{user}", "{user}1", "{user}111", "{user}123", "{user}@123", "{user}_123", "{user}#123", "{user}@111", "{user}@2019", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "000000", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system", "huawei"}

func init() {

	SuccessHash = make(map[string]bool)

	CommonPortMap = make(map[string]int)
	CommonPortMap["ftp"] = 21
	CommonPortMap["ssh"] = 22
	CommonPortMap["oxid"] = 135
	CommonPortMap["smb"] = 445
	CommonPortMap["mssql"] = 1433
	CommonPortMap["mysql"] = 6379
	CommonPortMap["redis"] = 6379

}

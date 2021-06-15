package model

import (
	"sync"
	"time"
)

const (
	TIMEUNIT = 3
	TIMEOUT  = time.Duration(TIMEUNIT) * time.Second
)

var (
	CommonPortMap map[string]int
	SuccessHash   map[string]bool
	Mutex         sync.Mutex
)

func init() {

	SuccessHash = make(map[string]bool)

	CommonPortMap = make(map[string]int)
	CommonPortMap["FTP"] = 21
	CommonPortMap["SSH"] = 22
	CommonPortMap["OXID"] = 135
	CommonPortMap["SMB"] = 445
	CommonPortMap["MSSQL"] = 1433
	CommonPortMap["MYSQL"] = 6379
	CommonPortMap["REDIS"] = 6379

}

package model

import "time"

const (
	TIMEUNIT = 3
	TIMEOUT  = time.Duration(TIMEUNIT) * time.Second
)

var (
	CommonPortMap map[string]int
)

func init() {

	CommonPortMap = make(map[string]int)
	CommonPortMap["FTP"] = 21
	CommonPortMap["SSH"] = 22
	CommonPortMap["SMB"] = 445
	CommonPortMap["MSSQL"] = 1433
	CommonPortMap["REDIS"] = 6379

}

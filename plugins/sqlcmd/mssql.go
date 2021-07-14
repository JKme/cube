package sqlcmd

import (
	"cube/log"
	"cube/model"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func Mssql1Cmd(task model.SqlcmdTask) (result model.SqlcmdTaskResult) {
	result = model.SqlcmdTaskResult{SqlcmdTask: task, Result: "", Err: nil}

	dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", task.Ip,
		task.Port, task.User, task.Password, "tempdb")
	db, err := sql.Open("mssql", dataSourceName)
	defer db.Close()
	if err != nil {
		log.Errorf("Open connection failed:", err.Error())
	}
	return result
}

func Open(conn sql.DB) {
	value, err := conn.Prepare("select value_in_use from   sys.configurations where  name = 'xp_cmdshell'")
	if err != nil {
		log.Errorf("Prepare failed:", err.Error())
	}
	defer value.Close()

	row := value.QueryRow()
	//var somenumber int64
	var v int
	err = row.Scan(&v)
	if err != nil {
		log.Errorf("Query failed:", err.Error())
	}
	if v == 1 {
		fmt.Printf("xp_cmdshell Enabled\n")

	} else {
		fmt.Printf("Open xp_cmdshell...\n")
		stmt, err := conn.Prepare("EXEC sp_configure 'show advanced options', 1;RECONFIGURE;EXEC sp_configure 'xp_cmdshell', 1;RECONFIGURE;")
		if err != nil {
			//fmt.Println("Query Error", err)
			return
		}

		defer stmt.Close()
		stmt.Query()

	}
	return

}

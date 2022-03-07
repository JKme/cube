package sqlcmdmodule

import (
	"bytes"
	"github.com/olekukonko/tablewriter"
)

type ISqlcmd interface {
	SqlcmdName() string //设定名称
	SqlcmdPort() string
	SqlcmdExec() SqlcmdResult
	SqlcmdDesc() string
}

func (s *Sqlcmd) NewISqlcmd() ISqlcmd {
	switch s.Name {
	case "mysql":
		return &Mysql{s}
	case "ssh":
		return &CmdSsh{s}
	case "mssql1":
		return &Mssql1{s}
	case "mssql2":
		return &Mssql2{s}
	case "mssql3":
		return &Mssql3{s}
	case "mssql4":
		return &Mssql4{s}
	default:
		return nil
	}
}

func NewSqlcmd(s string) Sqlcmd {
	return Sqlcmd{
		Name: s,
	}
}

func GetSqlcmdPort(s string) string {
	c := NewSqlcmd(s)
	ic := c.NewISqlcmd()
	return ic.SqlcmdPort()
}

func GetSqlcmdDesc(s string) string {
	c := NewSqlcmd(s)
	ic := c.NewISqlcmd()
	return ic.SqlcmdDesc()
}

var SqlcmdKeys = []string{"ssh", "mysql", "mssql1", "mssql2", "mssql3", "mssql4"}

func SqlcmdHelpTable() string {
	buf := bytes.NewBufferString("")
	table := tablewriter.NewWriter(buf)
	table.SetHeader([]string{"Func", "Port", "Desc"})
	for _, k := range SqlcmdKeys {
		table.Append([]string{k, GetSqlcmdPort(k), GetSqlcmdDesc(k)})
		table.SetRowLine(true)
	}
	table.Render()
	return buf.String()
}

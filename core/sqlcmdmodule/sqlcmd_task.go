package sqlcmdmodule

import (
	"cube/core"
	"cube/gologger"
	"fmt"
)

func StartSqlcmd(opt *SqlcmdOption, globalopt *core.GlobalOption) {
	ip := opt.Ip
	var port string
	if len(opt.Port) == 0 {
		port = GetSqlcmdPort(opt.Name)
	} else {
		port = opt.Port
	}
	user := opt.User
	pass := opt.Password
	e := opt.Query
	sc := Sqlcmd{
		Ip:       ip,
		Port:     port,
		User:     user,
		Password: pass,
		Query:    e,
		Name:     opt.Name,
	}
	sp := sc.NewISqlcmd()
	fn := sp.SqlcmdExec()
	if len(fn.Result) > 0 {
		s := fmt.Sprintf("[->>>>>]: %s\n[->>>>>]: %s:%s\n", fn.Sqlcmd.Name, fn.Sqlcmd.Ip, fn.Sqlcmd.Port)
		s1 := fmt.Sprintf("[output]:\n%s", fn.Result)
		gologger.Info(s + s1)
	} else {
		gologger.Info(fn.Err)
	}
}

package sqlcmd

import (
	"cube/model"
	"fmt"
	"testing"
)

func TestSshCmd(t *testing.T) {
	task := model.SqlcmdTask{Ip: "11", Port: 22, User: "ubuntu", Password: "", SqlcmdPlugin: "SSH", Query: "w"}
	//fmt.Println(SshCmd(task))
	r := SshCmd(task)
	fmt.Println(r.Result)
}

func TestMssql3Cmd(t *testing.T) {
	task := model.SqlcmdTask{Ip: "172.16.157.163", User: "sa", Password: "123456aa", SqlcmdPlugin: "mssql-clr", Query: "tasklist"}
	//fmt.Println(SshCmd(task))
	r := MssqlClr(task)
	fmt.Println(r.Result)
}

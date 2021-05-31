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

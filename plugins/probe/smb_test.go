package probe

import (
	"cube/model"
	"fmt"
	"testing"
)

//172.20.10.21
//172.16.157.190

func TestSMB(t *testing.T) {

	task := model.ProbeTask{Ip: "172.16.157.190", Port: "445", ScanPlugin: "smb"}
	r := SmbProbe(task)
	//fmt.Printf("%v\n", r.Result)
	fmt.Println(r.Result)
}

func TestWinrmProbe(t *testing.T) {
	task := model.ProbeTask{Ip: "172.20.10.21", Port: "5985", ScanPlugin: "smb"}
	r := WinrmProbe(task)
	fmt.Println(r.Result)
}

func TestWmiProbe(t *testing.T) {
	task := model.ProbeTask{Ip: "172.20.10.21", Port: "135", ScanPlugin: "smb"}
	r := WmiProbe(task)
	fmt.Println(r.Result)
}

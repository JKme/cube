package probe

import (
	"cube/model"
	"fmt"
	"testing"
)

func TestSMB(t *testing.T) {
	task := model.ProbeTask{Ip: "172.20.40.130", Port: "445", ScanPlugin: "SMB"}
	r := smbScan2(task)
	fmt.Printf("Get Result: %v\n", r.Result)
}

func TestWMI(t *testing.T) {
	task := model.ProbeTask{Ip: "172.20.40.124", Port: "135", ScanPlugin: "WMI"}
	wmi(task)
	//fmt.Printf("Get Result: %v\n", r.Result)
}

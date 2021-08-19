package probe

import (
	"cube/model"
	"fmt"
	"testing"
)

func TestSmbProbeV2(t *testing.T) {

	task := model.ProbeTask{Ip: "172.16.47.4", Port: "445", ScanPlugin: "smb"}
	r := SmbProbeV2(task)
	//fmt.Printf("%v\n", r.Result)
	fmt.Println(r.Result)
}

func TestSMB(t *testing.T) {

	task := model.ProbeTask{Ip: "172.16.47.4", Port: "445", ScanPlugin: "smb"}
	r := SmbProbe(task)
	//fmt.Printf("%v\n", r.Result)
	fmt.Println(r.Result)
}

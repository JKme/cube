package probe

import (
	"cube/model"
	"fmt"
	"testing"
)

//172.20.10.21
//172.16.157.190
func TestName(t *testing.T) {
	task := model.ProbeTask{Ip: "192.168.2.148", Port: "135", ScanPlugin: "oxid"}
	r := OxidProbe(task)
	//fmt.Println(reflect.TypeOf(r.Result))
	fmt.Println(r)

	//for _, v := range r.Result {
	//	fmt.Println( v)
	//	//r1.append
	//}
}

func TestZookeeperProbe(t *testing.T) {
	task := model.ProbeTask{Ip: "172.20.40.30", Port: "2182", ScanPlugin: "zookeeper"}
	r := ZookeeperProbe(task)
	//fmt.Println(reflect.TypeOf(r.Result))
	fmt.Println(r.Result)

	//for _, v := range r.Result {
	//	fmt.Println( v)
	//	//r1.append
	//}
}
func TestSMB(t *testing.T) {

	task := model.ProbeTask{Ip: "172.20.40.140", Port: "139", ScanPlugin: "smb"}
	r := SmbProbe(task)
	//fmt.Printf("%v\n", r.Result)
	fmt.Println(r.Result)
}

func TestWinrmProbe(t *testing.T) {
	task := model.ProbeTask{Ip: "192.168.2.226", Port: "5985", ScanPlugin: "ntlm-winrm"}
	r := WinrmProbe(task)
	fmt.Println(r.Result)
}

func TestWmiProbe(t *testing.T) {
	task := model.ProbeTask{Ip: "192.168.2.226", Port: "135", ScanPlugin: "smb"}
	r := WmiProbe(task)
	fmt.Println(r.Result)
}

func TestMssqlProbe(t *testing.T) {
	task := model.ProbeTask{Ip: "172.16.157.190", Port: "1433", ScanPlugin: "smb"}
	r := MssqlProbe(task)
	fmt.Println(r.Result)
}

func TestNetbiosProbe(t *testing.T) {
	task := model.ProbeTask{Ip: "172.20.40.140", Port: "137", ScanPlugin: "netbios"}
	r := NetbiosProbe(task)
	fmt.Println(r.Result)
}

func TestMdnsProbe(t *testing.T) {
	task := model.ProbeTask{Ip: "127.0.0.1", Port: "1099", ScanPlugin: "netbios"}
	r := RmiProbe(task)
	fmt.Println(r)
}

func TestDockerProbe(t *testing.T) {
	task := model.ProbeTask{Ip: "127.0.0.1", Port: "2375", ScanPlugin: "netbios"}
	r := DockerProbe(task)
	fmt.Println(r)
}

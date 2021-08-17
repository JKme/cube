package probe

import (
	"cube/model"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	task := model.ProbeTask{Ip: "172.20.40.108", Port: "135", ScanPlugin: "oxid"}
	r := OxidProbe(task)
	//fmt.Println(reflect.TypeOf(r.Result))
	fmt.Println(r.Result)

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

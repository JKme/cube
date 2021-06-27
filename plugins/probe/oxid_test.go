package probe

import (
	"cube/model"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	task := model.ProbeTask{Ip: " 192.168.2.226", Port: "135", ScanPlugin: "oxid"}
	r := OxidProbe(task)
	//fmt.Println(reflect.TypeOf(r.Result))
	fmt.Println(r.Result)

	//for _, v := range r.Result {
	//	fmt.Println( v)
	//	//r1.append
	//}
}

package probe

import (
	"cube/model"
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	task := model.ProbeTask{Ip: "172.20.40.100", Port: 135, ScanPlugin: "OXID"}
	r := OxidProbe(task)
	fmt.Println(reflect.TypeOf(r.Result))
	fmt.Printf("Get Result: %v\n", r)

	//for _, v := range r.Result {
	//	fmt.Println( v)
	//	//r1.append
	//}
}

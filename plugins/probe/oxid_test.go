package probe

import (
	"cube/model"
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	task := model.Task{Ip: "192.168.2.226 ", Port: 135, ScanPlugin: "OXID"}
	r := OxidScan(task)
	fmt.Println(reflect.TypeOf(r.Result))
	fmt.Printf("Get Result: %v\n", r.Result)

	//for _, v := range r.Result {
	//	fmt.Println( v)
	//	//r1.append
	//}
}

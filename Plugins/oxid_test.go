package Plugins

import (
	"NTLM/model"
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	task := model.Task{Ip: "172.20.42.5", Port: 135, ScanPlugin: "OXID"}
	r:= oxidScan(task)
	fmt.Println(reflect.TypeOf(r.Result))
	fmt.Printf("Get Result: %v\n", r.Result)

	//for _, v := range r.Result {
	//	fmt.Println( v)
	//	//r1.append
	//}
}
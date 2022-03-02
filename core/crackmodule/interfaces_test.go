package crackmodule

import (
	"cube/gologger"
	"fmt"
	"strconv"
	"testing"
)

func TestCrackPluginInterface(t *testing.T) {
	c := Crack{
		Ip:   "127.0.0.1",
		Port: "6379",
		Auth: Auth{
			User:     "",
			Password: "root",
		},
		Name: "redis",
	}

	task := c.NewICrack()
	if task == nil {
		gologger.Error("未找到插件")
	}
	task.Exec()

}

func TestParsePluginOpt(t *testing.T) {
	//l := ParsePluginOpt("smb")
	//fmt.Println(l)
	b, _ := strconv.ParseBool("222")
	fmt.Println(b)
}

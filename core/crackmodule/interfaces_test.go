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
		Port: "22",
		Auth: Auth{
			User:     "root",
			Password: "root",
		},
		Name: "ssh2",
	}

	fmt.Println(CrackKeys)
	task := c.NewICrack()
	if task == nil {
		gologger.Error("未找到插件")
	}

	gologger.Infof("INFO\n")

}

func TestParsePluginOpt(t *testing.T) {
	//l := ParsePluginOpt("smb")
	//fmt.Println(l)
	b, _ := strconv.ParseBool("222")
	fmt.Println(b)
}

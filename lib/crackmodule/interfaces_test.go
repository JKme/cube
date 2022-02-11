package crackmodule

import (
	"cube/lib/crackmodule/plugins"
	"fmt"
	"testing"
)


func Hello(c CrackPluginInterface){
	fmt.Println(c.Desc())
}

func TestCrackPluginInterface(t *testing.T) {
	s := plugins.SshCrack{
		Ip:   "127.0.0.1",
		Port: "22",
		Auth: plugins.Auth{
			User:     "root",
			Password: "123456",
		},
		Name: "ssh",
	}
	plugins.Crack{
		Ip:   "",
		Port: "",
		Auth: plugins.Auth{},
		Name: "",
	}
	Hello(s)
	//var c CrackPluginInterface
	//fmt.Println(s.Name())
	//fmt.Println(c.Port())
	//fmt.Println(c.Desc())
}
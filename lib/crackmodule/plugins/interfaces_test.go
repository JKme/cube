package plugins

import (
	"fmt"
	"testing"
)


func Hello(c ICrack){
	fmt.Println(c.Desc())
	fmt.Println(c.Exec())
}

func TestCrackPluginInterface(t *testing.T) {
	c := Crack{
		Ip:   "127.0.0.1",
		Port: "22",
		Auth: Auth{
			User:     "root",
			Password: "root",
		},
		Name: "ssh",
	}

	task := c.New()
	fmt.Println(task.Exec())
	fmt.Println(task.GetPort())
	fmt.Println(task.SetName())
	fmt.Println(task.Desc())
	fmt.Println(task.Load())
}

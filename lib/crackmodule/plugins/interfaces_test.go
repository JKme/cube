package plugins

import (
	"cube/gologger"
	"fmt"
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

	task := c.NewCrack()
	fmt.Println(task.Exec())
	fmt.Println(task.GetPort())
	fmt.Println(task.SetName())
	fmt.Println(task.Desc())
	fmt.Println(task.Load())
	gologger.Infof("INFO\n")

	gologger.Printf("Hello\n")
	gologger.Labelf("Label\n")
	gologger.Silentf("silent\n")
	gologger.Debugf("Debug\n")
	gologger.Fatalf("Fetalf\n")

}

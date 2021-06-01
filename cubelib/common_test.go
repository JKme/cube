package cubelib

import (
	"fmt"
	"testing"
)

//func TestName(t *testing.T) {
//	ParseService("127.0.0.1:22")
//}

func TestParseService(t *testing.T) {
	a, err := ParseService("ssh://127.0.0.11")

	fmt.Println(a, err)
}

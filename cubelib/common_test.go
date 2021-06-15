package cubelib

import (
	"fmt"
	"strings"
	"testing"
)

//func TestName(t *testing.T) {
//	ParseService("127.0.0.1:22")
//}

func TestParseService(t *testing.T) {
	a, err := ParseService("ssh://127.0.0.11")

	fmt.Println(a, err)
}

func TestSliceContain(t *testing.T) {
	s := "SMB,SSH"
	s1 := strings.Split(s, ",")
	fmt.Println(SliceContain("SMB2", s1))
}

func TestSameStringSlice(t *testing.T) {
	s := []string{"SMB", "SSH", "OXID"}
	s1 := []string{"SMB", "ERR"}
	//fmt.Println(SameStringSlice(s, s1))
	fmt.Println(Subset(s1, s))
}

func TestColor(t *testing.T) {
	fmt.Println(Fata("hello, world!"))
}

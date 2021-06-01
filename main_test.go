package main

import (
	"fmt"
	"strings"
	"testing"
)

type A struct {
	User     string
	Password string
}

type B = A

func TestName(t *testing.T) {
	a := A{User: "root", Password: "root"}
	fmt.Println(a)
	b := B{User: "root", Password: "111111"}
	fmt.Println(b.Password)

	s := "SMB,SSH"
	s1 := strings.Split(s, ",")
	fmt.Println(s1)
	for _, value := range s1 {
		if "SMB" == value {
			fmt.Println("success")
			break
		} else {
			fmt.Println("fail")
		}

	}
	//_, key := s1["SSH"]
	//x = []string{"SMB", "ALL"}

}

func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}

func TestsameStringSlice(t *testing.T) {

}

package main

import (
	"fmt"
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
}

package Plugins

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	for k, _ := range ProbeFuncMap {
		fmt.Println(k)
	}

	if _, ok := ProbeFuncMap["OXID"]; ok {
		fmt.Println(ok)
		//存在
	}

	//fmt.Println(ProbeKeys)
}

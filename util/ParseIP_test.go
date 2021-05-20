package util

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	ips, _ := ParseIP("172.20.42.5/24", "")
	fmt.Println(ips)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

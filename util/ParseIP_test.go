package util

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	ips, _ := ParseIP("172.20.42.5/24,10.0.0.5-a", "")
	fmt.Println(ips)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func TestIsUpper(t *testing.T) {
	r := IsUpper("ALL")
	fmt.Println(r)
}

//func TestCheckAlive(t *testing.T) {
//	ips, _ := ParseIP("172.20.40.22/24", "")
//	plugins := []string{"ssh"}
//	//r := CheckAlive(ips, plugins, "")
//	fmt.Println(r)
//}

func TestStrXor(t *testing.T) {
	r := StrXor("", "1")
	fmt.Println(r)
	print(r)

}

func TestReadipfile(t *testing.T) {
	r, _ := Readipfile("/tmp/ip.txt")
	fmt.Println(r)
}

func TestReadipfile2(t *testing.T) {
	ip := IpAddr{
		Ip:     "172.20.40.1",
		Port:   "137",
		Plugin: "netbios",
	}
	checkUDP(ip)

}

func TestSubset(t *testing.T) {
	a := []string{"ntlm-winrm"}
	b := []string{"ntlm-smb", "ntlm-wmi", "zookeeper", "oxid", "netbios", "ntlm-winrm"}
	fmt.Println(Subset(a, b))
}

func TestBytes2Uint(t *testing.T) {
	//a := []byte("\xcb\xef\xd2\xc0\xc1\xd5\x2d\x34\x35\x36\x20\x20\x20\x20\x20\x20")
	//s := bytes2StringUTF16(a)
	//fmt.Println(s)
}

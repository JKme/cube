package cubelib

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)


func ValidIp(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}

func ParseService(str string)(service string, ip string, port string){
	//ftp://192.168.0.1
	s := strings.Split(str, ":")

	ip, port = s[0], s[1]
	fmt.Println(ip, port)
	return service, ip, port
}

func Split(r rune) bool {
	return r == ':' || r == '://'
}

func ParseNet(str string)(service string, ip string, port string) {
	a := strings.FieldsFunc(str, Split)
}

// Split https://stackoverflow.com/questions/39862613/how-to-split-a-string-by-multiple-delimiters



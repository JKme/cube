package cubelib

import (
	"cube/model"
	"fmt"
	"regexp"
	"strconv"
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

func Split(r rune) bool {
	return strings.ContainsRune("://:", r)
}

func ParseService(str string) (service model.Service, err error) {
	a := strings.FieldsFunc(str, Split)
	l := len(a)
	if l < 2 || l > 3 {
		return service, fmt.Errorf("invalid target: %s (eg: ssh://192.168.1.1:22)", str)
	}

	service.Schema = strings.ToUpper(a[0])
	service.Ip = a[1]
	if !ValidIp(service.Ip) {
		return service, fmt.Errorf("invalid ip: %s", service.Ip)
	}

	if len(a) == 2 {
		service.Port = model.CommonPortMap[service.Schema]
	} else {
		service.Port, _ = strconv.Atoi(a[2])
	}

	return service, nil
}

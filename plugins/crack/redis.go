package crack

import (
	"cube/model"
	"cube/util"
	"fmt"
	"net"
	"regexp"
	"strings"
)

func RedisCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", task.Ip, task.Port), model.ConnectTimeout)
	if err != nil {
		return
	}

	_, err = conn.Write([]byte(fmt.Sprintf("auth %s\r\n", task.Auth.Password)))
	if err != nil {
		return
	}
	buf := make([]byte, 4096)
	count, _ := conn.Read(buf)
	config, _ := getConfig(conn)
	fmt.Printf("Config: %s\n", config)
	response := string(buf[0:count])
	if strings.Contains(response, "+OK") {
		result.Result = util.Green(fmt.Sprintf("Password: %s", task.Auth.Password))

	}

	return result
}
func readReply(conn net.Conn) (result string, err error) {
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			break
		}
		result += string(buf[0:count])
		if count < 4096 {
			break
		}
	}
	return result, err
}

func getConfig(conn net.Conn) (conf []string, err error) {
	_, err = conn.Write([]byte(fmt.Sprintf("INFO\r\n")))
	if err != nil {
		return
	}
	text, err := readReply(conn)
	if err != nil {
		return
	}
	//l := strings.Split(text, "\n")
	fmt.Println(strings.Split(text, "\n")[2])
	r := regexp.MustCompile(`.*redis_version:(.*)\n(?s).*(?U)os:(.*)\n`)
	l := r.FindStringSubmatch(text)
	fmt.Println(l[1], l[2])
	fmt.Printf("[+] %#v\n", r.FindStringSubmatch(text))

	//text1 := strings.Split(text, "\n")
	//if len(text1) > 2 {
	//	dbfilename = text1[len(text1)-2]
	//} else {
	//	dbfilename = text1[0]
	//}
	return
}

//redis_version  os

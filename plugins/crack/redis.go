package crack

import (
	"cube/log"
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
	log.Debugf("Password: %s", task.Auth.Password)

	config, err := getConfig(conn)
	if err != nil {
		return
	}

	if len(config) > 0 {
		result.Result = util.Green(fmt.Sprintf("Password: %s \t Version=%s  OS=%s", task.Auth.Password, config[0], config[1]))
	} else {
		_, err = conn.Write([]byte(fmt.Sprintf("AUTH %s\r\n", task.Auth.Password)))
		if err != nil {
			return
		}
		buf := make([]byte, 4096)
		count, _ := conn.Read(buf)
		response := string(buf[0:count])
		if strings.Contains(response, "+OK") {
			config, _ := getConfig(conn)
			log.Debugf("Config: %v", config)
			//result.Result = util.Green(fmt.Sprintf("Password: %s \t Version: %v ", task.Auth.Password, config))

			result.Result = util.Green(fmt.Sprintf("Password: %s \t Version=%s  OS=%s", task.Auth.Password, config[0], config[1]))
		}
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
	if strings.Contains(text, "redis_version") {
		r := regexp.MustCompile(`.*redis_version:(.*)\n(?s).*(?U)os:(.*)\n`)

		match := r.FindStringSubmatch(text)
		a := strings.TrimSpace(match[1])
		b := strings.TrimSpace(match[2])
		conf = append(conf, a, b)
	}
	return conf, nil
}

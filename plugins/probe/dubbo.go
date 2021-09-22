package probe

import (
	"cube/model"
	"cube/util"
	"fmt"
	"net"
)

func DubboProbe(task model.ProbeTask) (result model.ProbeTaskResult) {
	result = model.ProbeTaskResult{ProbeTask: task, Result: "", Err: nil}
	host := fmt.Sprintf("%s:%v", task.Ip, task.Port)
	conn, _ := net.DialTimeout("tcp", host, model.ConnectTimeout)
	_, err := conn.Write([]byte("\r\n\r\n"))
	if err != nil {
		return
	}
	r1, _ := util.ReadBytes(conn)
	//fmt.Printf("Receive: %s\n", string(r1[:5]))
	if string(r1[:5]) == "dubbo" {
		result.Result = "Dubbo Service"
	}

	return result
}

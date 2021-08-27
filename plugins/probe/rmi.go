package probe

import (
	"cube/model"
	"cube/util"
	"encoding/hex"
	"fmt"
	"net"
)

// https://koalr.me/post/fastjson-deserialization-detection/
func RmiProbe(task model.ProbeTask) (result model.ProbeTaskResult) {
	result = model.ProbeTaskResult{ProbeTask: task, Result: "", Err: nil}
	host := fmt.Sprintf("%s:%v", task.Ip, task.Port)
	conn, _ := net.DialTimeout("tcp", host, model.ConnectTimeout)
	_, err := conn.Write([]byte{0x4a, 0x52, 0x4d, 0x49, 0x00, 0x02, 0x4b})
	if err != nil {
		return
	}
	r1, _ := util.ReadBytes(conn)
	//fmt.Printf("%x", r1[:1])
	if hex.EncodeToString(r1[:1]) == "4e" {
		result.Result = "RMI Registry Deserialization"
	}

	return result
}

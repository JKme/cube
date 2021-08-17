package probe

import (
	"cube/model"
	"fmt"
	"net"
	"strings"
)

func ZookeeperProbe(task model.ProbeTask) (result model.ProbeTaskResult) {
	result = model.ProbeTaskResult{ProbeTask: task, Result: "", Err: nil}
	dl := net.Dialer{Timeout: model.ConnectTimeout}
	t := fmt.Sprintf("%s:%d", task.Ip, model.CommonPortMap[task.ScanPlugin])
	conn, err := dl.Dial("tcp", t)
	if err != nil {
		result.Err = err
		//log.Printf("Oxid Running Error: %s:%s", task.Ip, err)
		return result
	}
	conn.Write([]byte("version\r\n"))

	r := make([]byte, 1024)
	conn.Read(r)
	if strings.Contains(string(r), "ZooKeeper") {
		result.Result = "Zookeeper Unauthorized"
	}
	return result
}

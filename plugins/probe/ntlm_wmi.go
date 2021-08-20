package probe

import (
	"bytes"
	"cube/log"
	"cube/model"
	"fmt"
	"github.com/JKme/go-ntlmssp"
	"net"
)

var payload = []byte{5, 0, 11, 3, 16, 0, 0, 0, 120, 0, 40, 0, 3, 0, 0, 0, 184, 16, 184, 16, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 160, 1, 0, 0, 0, 0, 0, 0, 192, 0, 0, 0, 0, 0, 0, 70, 0, 0, 0, 0, 4, 93, 136, 138, 235, 28, 201, 17, 159, 232, 8, 0, 43, 16, 72, 96, 2, 0, 0, 0, 10, 2, 0, 0, 0, 0, 0, 0, 78, 84, 76, 77, 83, 83, 80, 0, 1, 0, 0, 0, 7, 130, 8, 162, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 1, 177, 29, 0, 0, 0, 15}

func WmiProbe(task model.ProbeTask) (result model.ProbeTaskResult) {
	result = model.ProbeTaskResult{ProbeTask: task, Result: "", Err: nil}
	realhost := fmt.Sprintf("%s:%v", task.Ip, task.Port)
	conn, err := net.DialTimeout("tcp", realhost, model.ConnectTimeout)
	if err != nil {
		log.Debug(err)
		return
	}
	conn.Write(payload)
	if err != nil {
		return
	}
	ret, _ := readBytes(conn)

	off_ntlm := bytes.Index(ret, []byte("NTLMSSP"))
	type2 := ntlmssp.ChallengeMsg{}
	tinfo := "\n" + type2.String(ret[off_ntlm:])
	result.Result = tinfo
	return result
}

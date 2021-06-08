package crack

import (
	"cube/model"
	"time"
)

func SshCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	time.Sleep(1 * time.Second)
	result.Result = "[*]: SSH Crack Plugin Done"
	return result

}

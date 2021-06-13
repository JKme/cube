package crack

import (
	"cube/model"
	"fmt"
	"time"
)

func SshCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	fmt.Printf("Receive Task: %s\n", task)
	time.Sleep(1 * time.Second)
	if task.Auth.Password == "111" {
		result.Result = "[*]: SSH Crack Plugin Done"
	}
	return result

}

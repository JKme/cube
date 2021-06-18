package crack

import (
	"cube/model"
	"fmt"
)

func SshCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	fmt.Printf("Receive Task: %s\n", task)
	//time.Sleep(2 * time.Second)
	if task.Auth.Password == "root1" {
		result.Result = "SSH Crack Plugin Done"
	}
	return result

}

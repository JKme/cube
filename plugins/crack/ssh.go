package crack

import "cube/model"

func SshCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	result.Result = "Test"
	return result

}

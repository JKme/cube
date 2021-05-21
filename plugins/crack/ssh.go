package crack

import "cube/model"

func SshCrack(task model.Task) (result model.TaskResult) {
	result = model.TaskResult{Task: task, Result: "", Err: nil}

	return result

}

package sqlcmd

import "cube/model"

func SshCmd(task model.SqlcmdTask) (result model.SqlcmdTaskResult) {
	result = model.SqlcmdTaskResult{SqlcmdTask: task, Result: "", Err: nil}

	return result

}

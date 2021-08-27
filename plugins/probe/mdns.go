package probe

import (
	"cube/model"
)

func MdnsProbe(task model.ProbeTask) (result model.ProbeTaskResult) {
	result = model.ProbeTaskResult{ProbeTask: task, Result: "", Err: nil}

	return result
}

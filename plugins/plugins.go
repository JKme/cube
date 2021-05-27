package Plugins

import (
	"cube/model"
	"cube/plugins/probe"
)

type ProbeFunc func(task model.ProbeTask) (taskResult model.ProbeTaskResult)

//type CrackFunc func(task model.Task) (taskResult model.TaskResult)

var (
	ProbeFuncMap map[string]ProbeFunc
	//CrackFuncMap map[string]CrackFunc
)

func init() {
	ProbeFuncMap = make(map[string]ProbeFunc)
	ProbeFuncMap["OXID"] = probe.OxidProbe
	ProbeFuncMap["SMB"] = probe.SmbProbe

	//CrackFuncMap = make(map[string]CrackFunc)
	//CrackFuncMap["SSH"] = crack.SshCrack

}

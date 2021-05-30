package Plugins

import (
	"cube/model"
	"cube/plugins/probe"
)

type ProbeFunc func(task model.ProbeTask) (taskResult model.ProbeTaskResult)

//type CrackFunc func(task model.Task) (taskResult model.TaskResult)

var (
	ProbeFuncMap map[string]ProbeFunc
	ProbeKeys    []string
	//CrackFuncMap map[string]CrackFunc
)

func init() {
	ProbeFuncMap = make(map[string]ProbeFunc)
	ProbeFuncMap["OXID"] = probe.OxidProbe
	ProbeFuncMap["SMB"] = probe.SmbProbe

	ProbeKeys = append(ProbeKeys, "ALL")
	for k := range ProbeFuncMap {
		ProbeKeys = append(ProbeKeys, k)
	}
	//CrackFuncMap = make(map[string]CrackFunc)
	//CrackFuncMap["SSH"] = crack.SshCrack
	//var keys []string
	//for k := range ProbeFuncMap {
	//	keys = append(keys, k)
	//}
}

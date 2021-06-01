package Plugins

import (
	"cube/model"
	"cube/plugins/probe"
	"cube/plugins/sqlcmd"
)

type ProbeFunc func(task model.ProbeTask) (taskResult model.ProbeTaskResult)
type SqlcmdFunc func(task model.SqlcmdTask) (taskResult model.SqlcmdTaskResult)

//type CrackFunc func(task model.Task) (taskResult model.TaskResult)

var (
	ProbeFuncMap  map[string]ProbeFunc
	SqlcmdFuncMap map[string]SqlcmdFunc

	ProbeKeys  []string
	SqlcmdKeys []string
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

	SqlcmdFuncMap = make(map[string]SqlcmdFunc)
	SqlcmdFuncMap["SSH"] = sqlcmd.SshCmd
	for k := range SqlcmdFuncMap {
		SqlcmdKeys = append(SqlcmdKeys, k)
	}

	//CrackFuncMap = make(map[string]CrackFunc)
	//CrackFuncMap["SSH"] = crack.SshCrack
	//var keys []string
	//for k := range ProbeFuncMap {
	//	keys = append(keys, k)
	//}
}

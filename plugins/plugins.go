package Plugins

import (
	"cube/model"
)

type ScanFunc func(task model.Task) (taskResult model.TaskResult)

var (
	ScanFuncMap map[string]ScanFunc
)

func init(){
	ScanFuncMap = make(map[string]ScanFunc)
	ScanFuncMap["OXID"] = oxidScan
	ScanFuncMap["SMB"]  = smbScan
}
package plugins

import (
	"cube/pkg/util"
	"strings"
)

func ParsePluginOpt(n string) (pluginNameList []string) {
	//
	pns := strings.Split(n, ",")
	if len(pns) > 2 && util.Contains("X", pns) {
		//指定-X只能单独加载
		pluginNameList = nil
	}
	if len(pns) > 2 && util.Contains("Y", pns) {
		pluginNameList = nil
	}
	switch {
	case len(pns) == 1:
		if pns[0] == "X" {
			for _, k := range CrackKeys {
				if !GetMutexStatus(k) && GetLoadStatus(k) == "Y" {
					pluginNameList = append(pluginNameList, k)
				}
			}
		}
		if pns[0] == "Y" {
			for _, k := range CrackKeys {
				if !GetMutexStatus(k) {
					pluginNameList = append(pluginNameList, k)
				}
			}
		}
		if util.Contains(pns[0], CrackKeys) {
			pluginNameList = pns
		}
	default:
		for _, k := range pns {
			if util.Contains(k, CrackKeys) {
				pluginNameList = strings.Split(n, ",")
			} else {
				pluginNameList = nil
			}
		}

	}
	return pluginNameList
}

//func checkPluginName(name string) bool{
//	pluginList := strings.Split(name, ",")
//	for p, _ := range pluginList{
//		if GetMutexStatus(p){
//
//		}
//	}
//}

package plugins

import (
	"cube/core"
	"cube/pkg/util"
)

func buildTasks() {
	// 生成任务
}

func saveResult() {

}

func runSingleTask() {

}

func parseAuthOption() {

}

func parsePluginOption() {

}

func getAuthList() {

}

func StartCrack(opt *CrackOptions, globalopts *core.GlobalOptions) {
	var (
		optPlugins []string
		ips        []string
		auths      []Auth
		crackTasks []Crack
		num        int
		delay      int
		aliveIPS   []util.IpAddr
	)
}

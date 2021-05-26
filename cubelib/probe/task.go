package probe

import "cube/model"

func generateTasks(ipList []string, port int, scanPlugin string) (tasks []model.ProbeTask) {
	tasks = make([]model.ProbeTask, 0)
	for _, ip := range ipList {
		service := model.ProbeTask{Ip: ip, Port: port, ScanPlugin: scanPlugin}
		tasks = append(tasks, service)
	}
	return tasks
}

func getConfig(options model.ProbeOptions, globalOptions model.GlobalOptions) {

}

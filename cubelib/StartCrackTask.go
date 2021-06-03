package cubelib

import "cube/model"

func GenerateCrackTasks(ip []string, port string, user []string, password []string, plugins []string) (tasks []model.CrackTask) {
	tasks = make([]model.CrackTask, 0)
	for _, i := range ip {
		for _, u := range user {
			for _, pass := range password {
				for _, p := range plugins {
					s := model.CrackTask{Ip: i, Port: port, User: u, Password: pass, CrackPlugin: p}
					tasks = append(tasks, s)
				}
			}
		}
	}
	return tasks
}

func processArgs(args string) ([]string, error) {

	return nil, nil
}

func StartCrackTask(opt *model.CrackOptions, globalopts *model.GlobalOptions) {

}

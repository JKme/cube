package cubelib

import "cube/model"

func GenerateCrackTasks(ip []string, port string, auths []model.Auth, plugins []string) (tasks []model.CrackTask) {
	tasks = make([]model.CrackTask, 0)
	for _, i := range ip {
		for _, auth := range auths {
			for _, p := range plugins {
				s := model.CrackTask{Ip: i, Port: port, Auth: &auth, CrackPlugin: p}
				tasks = append(tasks, s)
			}
		}
	}
	return tasks
}

func processArgs(opt *model.CrackOptions) ([]string, error) {

	return nil, nil
}

func generateAuth(user []string, password []string) (authList []model.Auth) {
	authList = make([]model.Auth, 0)
	for _, u := range user {
		for _, pass := range password {
			a := model.Auth{User: u, Password: pass}
			authList = append(authList, a)
		}
	}
	return authList
}

func StartCrackTask(opt *model.CrackOptions, globalopts *model.GlobalOptions) {

}

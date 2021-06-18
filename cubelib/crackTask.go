package cubelib

import (
	"context"
	"cube/log"
	"cube/model"
	Plugins "cube/plugins"
	"fmt"
	"strings"
	"sync"
	"time"
)

func loadDefaultDict(p string) map[string][]model.Auth {
	authSlice := make([]model.Auth, 0)
	r := make(map[string][]model.Auth, 0)
	for _, user := range model.UserDict[p] {
		for _, pass := range model.PassDict {
			pass = strings.Replace(pass, "{user}", user, -1)
			authSlice = append(authSlice, model.Auth{
				User:     user,
				Password: pass,
			})
		}
	}
	r[p] = authSlice
	return r
}

func genDefaultTasks(ips []string, plugins []string) (tasks []model.CrackTask) {
	tasks = make([]model.CrackTask, 0)
	for _, ip := range ips {
		for _, plugin := range plugins {
			mapAuthSlice := loadDefaultDict(plugin)
			authSlice := mapAuthSlice[plugin]
			for _, auth := range authSlice {
				s := model.CrackTask{Ip: ip, Auth: auth, CrackPlugin: plugin}
				tasks = append(tasks, s)
			}

		}

	}
	return tasks
}

func genCrackTasks(plugins []string, ips []string, auths []model.Auth) (tasks []model.CrackTask) {
	tasks = make([]model.CrackTask, 0)
	for _, ip := range ips {
		for _, auth := range auths {
			for _, p := range plugins {
				s := model.CrackTask{Ip: ip, Auth: auth, CrackPlugin: p}
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

func saveCrackReport(taskResult model.CrackTaskResult) {

	if len(taskResult.Result) > 0 {
		k := fmt.Sprintf("%v-%v-%v", taskResult.CrackTask.Ip, taskResult.CrackTask.Port, taskResult.CrackTask.CrackPlugin)
		h := MakeTaskHash(k)
		SetTaskHash(h)
		s := fmt.Sprintf("[->>>>>]: %s\n[->>>>>]: %s:%s", taskResult.CrackTask.CrackPlugin, taskResult.CrackTask.Ip, taskResult.CrackTask.Port)
		s1 := fmt.Sprintf("[output]: %s", taskResult.Result)
		fmt.Println(s)
		fmt.Println(Green(s1))
	}
}

func runUnitTask(ctx context.Context, tasks chan model.CrackTask, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-tasks:
			if !ok {
				return
			}
			k := fmt.Sprintf("%v-%v-%v", task.Ip, task.Port, task.CrackPlugin)
			h := MakeTaskHash(k)
			if CheckTashHash(h) {
				wg.Done()
				continue
			}
			fn := Plugins.CrackFuncMap[task.CrackPlugin]
			r := fn(task)
			saveCrackReport(r)
			wg.Done()
			if len(r.Result) > 0 {
				log.Info(r)
			}

			select {
			case <-ctx.Done():
			case <-time.After(2 * time.Second):
			}

		}

	}
}

func runCrack(ctx context.Context, tasks []model.CrackTask) {

	var wg sync.WaitGroup
	taskChan := make(chan model.CrackTask, 8)

	for i := 0; i < 1; i++ {
		go runUnitTask(ctx, taskChan, &wg)
	}

	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}
	waitTimeout(&wg, model.T)
}

func opt2slice(str string, file string) []string {
	if len(str) > 0 {
		r := strings.Split(str, ",")
		return r
	}
	r, _ := FileReader(file)
	return r
}

func genPlugins(plugin string) []string {
	pluginList := strings.Split(plugin, ",")
	if len(pluginList) > 1 && SliceContain("ALL", pluginList) {
		log.Errorf("invalid plugin: %s", plugin)
	}

	if plugin == "ALL" {
		pluginList = Plugins.CrackKeys[1:]
	}
	return pluginList
}

func parseOpt(opt *model.CrackOptions) (ips []string, authList []model.Auth) {
	ip := opt.Ip
	ipFile := opt.IpFile

	ips = opt2slice(ip, ipFile)

	user := opt.User
	userFile := opt.UserFile
	pass := opt.Pass
	passFile := opt.PassFile
	us := opt2slice(user, userFile)
	ps := opt2slice(pass, passFile)

	for _, u := range us {
		for _, p := range ps {
			authList = append(authList, model.Auth{
				User:     u,
				Password: p,
			})
		}
	}

	return ips, authList
}

func StartCrackTask(opt *model.CrackOptions, globalopts *model.GlobalOptions) {
	var (
		optPlugins []string
		ips        []string
		auths      []model.Auth
		tasks      []model.CrackTask
	)
	optPlugins = genPlugins(opt.CrackPlugin)

	if len(opt.User+opt.UserFile+opt.Pass+opt.PassFile) > 0 {
		ips, auths = parseOpt(opt)
		tasks = genCrackTasks(optPlugins, ips, auths)
		log.Info(tasks)
	} else {
		tasks = genDefaultTasks(ips, optPlugins)
	}

	ctx := context.Background()

	var wg sync.WaitGroup
	taskChan := make(chan model.CrackTask, 8)

	for i := 0; i < 1; i++ {
		go runUnitTask(ctx, taskChan, &wg)
	}

	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}
	//wg.Wait()
	waitTimeout(&wg, model.T)
}

//func startCrackTask2(ips []string, authList []model.Auth, plugins []string) {
//	ctx := context.Background()
//	tasks := genCrackTasks(ips, authList, plugins)
//
//	var wg sync.WaitGroup
//	taskChan := make(chan model.CrackTask, 8)
//
//	for i := 0; i < 1; i++ {
//		go runUnitTask(ctx, taskChan, &wg)
//	}
//
//	for _, task := range tasks {
//		wg.Add(1)
//		taskChan <- task
//	}
//	//wg.Wait()
//	waitTimeout(&wg, model.T)
//}

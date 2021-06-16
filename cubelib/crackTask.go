package cubelib

import (
	"cube/log"
	"cube/model"
	Plugins "cube/plugins"
	"fmt"
	"strings"
	"sync"
	"time"
)

func unitTask(ips []string, auths []model.Auth, plugins []string) (tasks []model.CrackTask) {
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

func executeCrackTask(taskChan chan model.CrackTask, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskChan {
		//fmt.Println("Hello")
		fn := Plugins.CrackFuncMap[task.CrackPlugin]
		saveCrackReport(fn(task))
	}

}

func RunCrackTasks(tasks []model.CrackTask, scanNum int, timeout int) {
	tasksChan := make(chan model.CrackTask, scanNum*2)
	var wg sync.WaitGroup

	//消费者
	wg.Add(scanNum)
	for i := 0; i < scanNum; i++ {
		go executeCrackTask(tasksChan, &wg)
	}

	//生产者
	//go func() {
	//
	//}()

	for _, task := range tasks {
		tasksChan <- task
	}
	close(tasksChan)

	waitTimeout(&wg, time.Duration(timeout)*time.Second)
}

func saveCrackReport(taskResult model.CrackTaskResult) {

	if len(taskResult.Result) > 0 {
		k := fmt.Sprintf("%v-%v-%v", taskResult.CrackTask.Ip, taskResult.CrackTask.Port, taskResult.CrackTask.CrackPlugin)
		h := MakeTaskHash(k)
		SetTaskHash(h)
		s := fmt.Sprintf("[->>>>>]: %s\n[->>>>>]: %s:%s", taskResult.CrackTask.CrackPlugin, taskResult.CrackTask.Ip, taskResult.CrackTask.Port)
		s1 := fmt.Sprintf("[output]: %s", taskResult.Result)
		fmt.Println(s)
		fmt.Println(Fata(s1))
	}
}

func runUnitTask(tasks chan model.CrackTask, wg *sync.WaitGroup) {
	for task := range tasks {

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
			fmt.Println(r)
		}

	}
}

func runCrack(tasks []model.CrackTask) {

	var wg sync.WaitGroup
	taskChan := make(chan model.CrackTask, 8)

	for i := 0; i < 3; i++ {
		go runUnitTask(taskChan, &wg)
	}

	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}
	//wg.Wait()
	waitTimeout(&wg, model.TIMEOUT)
}

//func startCrackTask(plugins []string, ips []string, authList []model.Auth) {
//	tasks := unitTask(ips, authList, plugins)
//	runCrack(tasks)
//}

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
		log.Fatalf("invalid plugin: %s", plugin)
	}

	if plugin == "ALL" {
		pluginList = Plugins.CrackKeys[1:]
	}
	return pluginList
}

func parseOpt(opt *model.CrackOptions) (plugins []string, ips []string, authList []model.Auth) {
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

	plugin := opt.CrackPlugin
	plugins = genPlugins(plugin)

	return plugins, ips, authList
}

func startCrackTask(opt *model.CrackOptions, globalopts *model.GlobalOptions) {
	plugins, ips, authList := parseOpt(opt)
	tasks := unitTask(ips, authList, plugins)
	runCrack(tasks)
}

// https://stackoverflow.com/questions/45500836/close-multiple-goroutine-if-an-error-occurs-in-one-in-go

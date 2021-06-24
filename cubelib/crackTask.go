package cubelib

import (
	"context"
	"cube/log"
	"cube/model"
	Plugins "cube/plugins"
	"cube/util"
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

func genDefaultTasks(AliveAddr []util.IpAddr, plugins []string) (tasks []model.CrackTask) {
	tasks = make([]model.CrackTask, 0)
	for _, addr := range AliveAddr {
		for _, plugin := range plugins {
			mapAuthSlice := loadDefaultDict(plugin)
			authSlice := mapAuthSlice[plugin]
			for _, auth := range authSlice {
				s := model.CrackTask{Ip: addr.Ip, Port: addr.Port, Auth: auth, CrackPlugin: plugin}
				tasks = append(tasks, s)
			}

		}

	}
	return tasks
}

func genCrackTasks(plugins []string, AliveAddr []util.IpAddr, auths []model.Auth) (tasks []model.CrackTask) {
	tasks = make([]model.CrackTask, 0)
	for _, addr := range AliveAddr {
		for _, auth := range auths {
			for _, p := range plugins {
				s := model.CrackTask{Ip: addr.Ip, Port: addr.Port, Auth: auth, CrackPlugin: p}
				tasks = append(tasks, s)
			}
		}
	}
	return tasks
}

func processArgs(opt *model.CrackOptions) ([]string, error) {

	return nil, nil
}

func saveCrackReport(taskResult model.CrackTaskResult) {

	if len(taskResult.Result) > 0 {
		k := fmt.Sprintf("%v-%v-%v", taskResult.CrackTask.Ip, taskResult.CrackTask.Port, taskResult.CrackTask.CrackPlugin)
		h := MakeTaskHash(k)
		SetTaskHash(h)
		s1 := fmt.Sprintf("[+]: %s://%s:%s %s", taskResult.CrackTask.CrackPlugin, taskResult.CrackTask.Ip, taskResult.CrackTask.Port, taskResult.Result)
		fmt.Println(s1)
	}
}

func runUnitTask(ctx context.Context, tasks chan model.CrackTask, wg *sync.WaitGroup, delay int) {
	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-tasks:
			if !ok {
				return
			}
			log.Debugf("Checking %s Password: %s://%s:%s@%s:%s", task.CrackPlugin, task.CrackPlugin, task.Auth.User, task.Auth.Password, task.Ip, task.Port)
			k := fmt.Sprintf("%v-%v-%v", task.Ip, task.Port, task.CrackPlugin)
			h := MakeTaskHash(k)
			if CheckTaskHash(h) {
				wg.Done()
				continue
			}
			fn := Plugins.CrackFuncMap[task.CrackPlugin]
			r := fn(task)
			saveCrackReport(r)
			wg.Done()

			select {
			case <-ctx.Done():
			case <-time.After(time.Duration(delay) * time.Second):
			}

		}

	}
}

func opt2slice(str string, file string) []string {
	if len(str+file) == 0 {
		log.Error("-h for Help, Please set User and Password flag")
	}
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

func genAuths(opt *model.CrackOptions) (auths []model.Auth) {
	user := opt.User
	userFile := opt.UserFile
	pass := opt.Pass
	passFile := opt.PassFile
	us := opt2slice(user, userFile)
	ps := opt2slice(pass, passFile)

	for _, u := range us {
		for _, p := range ps {
			auths = append(auths, model.Auth{
				User:     u,
				Password: p,
			})
		}
	}
	return auths
}

func StartCrackTask(opt *model.CrackOptions, globalopts *model.GlobalOptions) {
	var (
		optPlugins []string
		ips        []string
		auths      []model.Auth
		tasks      []model.CrackTask
		num        int
		delay      int
	)
	ctx := context.Background()

	delay = globalopts.Delay

	if delay > 0 {
		num = 1
	} else {
		num = globalopts.Threads
	}

	optPlugins = genPlugins(opt.CrackPlugin)
	ips, _ = util.ParseIP(opt.Ip, opt.IpFile)
	AliveAddr := util.CheckAlive(ctx, num, delay, ips, optPlugins, opt.Port)

	if len(opt.User+opt.UserFile+opt.Pass+opt.PassFile) > 0 {
		auths = genAuths(opt)
		tasks = genCrackTasks(optPlugins, AliveAddr, auths)
	} else {
		tasks = genDefaultTasks(AliveAddr, optPlugins)
	}
	log.Debugf("Receive %d task", len(tasks))

	var wg sync.WaitGroup
	taskChan := make(chan model.CrackTask, num*2)

	for i := 0; i < num; i++ {
		go runUnitTask(ctx, taskChan, &wg, delay)

	}

	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task

	}
	//wg.Wait()
	waitTimeout(&wg, model.ThreadTimeout)

}

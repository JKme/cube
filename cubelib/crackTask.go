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
	users, ok := model.UserDict[p]
	if !ok {
		users = []string{""}
	}
	for _, user := range users {
		log.Debugf("User: %s", user)
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

func genDefaultTasks(AliveIPS []util.IpAddr) (tasks []model.CrackTask) {
	tasks = make([]model.CrackTask, 0)
	for _, addr := range AliveIPS {
		mapAuthSlice := loadDefaultDict(addr.Plugin)
		authSlice := mapAuthSlice[addr.Plugin]
		for _, auth := range authSlice {
			s := model.CrackTask{Ip: addr.Ip, Port: addr.Port, Auth: auth, CrackPlugin: addr.Plugin}
			tasks = append(tasks, s)
		}

	}
	return tasks
}

func genCrackTasks(AliveIPS []util.IpAddr, auths []model.Auth) (tasks []model.CrackTask) {
	tasks = make([]model.CrackTask, 0)
	for _, addr := range AliveIPS {
		for _, auth := range auths {
			s := model.CrackTask{Ip: addr.Ip, Port: addr.Port, Auth: auth, CrackPlugin: addr.Plugin}
			tasks = append(tasks, s)
		}
	}
	return tasks
}

func saveCrackReport(taskResult model.CrackTaskResult) {

	if len(taskResult.Result) > 0 {
		k := fmt.Sprintf("%v-%v-%v", taskResult.CrackTask.Ip, taskResult.CrackTask.Port, taskResult.CrackTask.CrackPlugin)
		h := MakeTaskHash(k)
		SetTaskHash(h)
		//s1 := fmt.Sprintf("[+]: %s://%s:%s %s", taskResult.CrackTask.CrackPlugin, taskResult.CrackTask.Ip, taskResult.CrackTask.Port, taskResult.Result)
		//fmt.Println(s1)
		SetResultMap(taskResult)
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
			//if task.Port == "" {
			//	task.Port =  strconv.Itoa(model.CommonPortMap[task.CrackPlugin])
			//}
			//alive := CheckAlive(task)
			//if !alive {
			//	wg.Done()
			//	continue
			//}

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
		pluginList = Plugins.CrackKeys
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
		AliveIPS   []util.IpAddr
	)
	ctx := context.Background()
	t1 := time.Now()
	delay = globalopts.Delay

	if delay > 0 {
		num = 1
	} else {
		num = globalopts.Threads
	}

	if opt.CrackPlugin == "phpmyadmin" {
		AliveIPS = append(AliveIPS, util.IpAddr{
			Ip:     opt.Ip,
			Port:   "",
			Plugin: opt.CrackPlugin,
		})
	} else {
		optPlugins = genPlugins(opt.CrackPlugin)
		log.Infof("Loading plugin: %s", strings.Join(optPlugins, ","))
		ips, _ = util.ParseIP(opt.Ip, opt.IpFile)

		AliveIPS = util.CheckAlive(ctx, num, delay, ips, optPlugins, opt.Port)
		//AliveIPS = RemoveRepByMap(AliveIPS)  // 去重IP
		log.Debugf("Receive alive IP: %s", AliveIPS)
	}

	if len(opt.User+opt.UserFile+opt.Pass+opt.PassFile) > 0 {
		auths = genAuths(opt)
		tasks = genCrackTasks(AliveIPS, auths)
	} else {
		tasks = genDefaultTasks(AliveIPS)
	}
	log.Debugf("Receive %d tasks", len(tasks))

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
	ReadResultMap()
	getFinishTime(t1)
}

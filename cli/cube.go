package cli

import (
	"cube/cubelib"
	"cube/log"
	"cube/model"
	Plugins "cube/plugins"
	"cube/util"
	"sync"
	"time"
)

func StartProbeTask(opt *model.ProbeOptions, globalopts *model.GlobalOptions) {
	ips, err := util.ParseIP(opt.Target, opt.TargetFile)
	if err != nil {
		log.Error(err)
	}
	tasks := cubelib.GenerateTasks(ips, opt.Port, opt.ScanPlugin)
	cubelib.RunTasks(tasks, globalopts.Threads, globalopts.Timeout)
}

func StartSqlcmdTask(opt *model.SqlcmdOptions, globalopts *model.GlobalOptions) {
	//TODO 前置判断条件，比如IP正则，Plugin
	//ips, err := util.ParseIP(opt.Target, opt.TargetFile)
	//if err != nil {
	//	log.Error(err)
	//}
	task := model.SqlcmdTask{Ip: opt.Ip, Port: opt.Port, User: opt.User, Password: opt.Password, SqlcmdPlugin: opt.SqlcmdPlugin, Query: opt.Query}
	fn := Plugins.SqlcmdFuncMap[task.SqlcmdPlugin]
	cubelib.SaveSqlcmdReport(fn(task))
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

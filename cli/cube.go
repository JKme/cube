package cli

import (
	"cube/cubelib"
	"cube/log"
	"cube/model"
	Plugins "cube/plugins"
	"cube/util"
	"strings"
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
	s, err := cubelib.ParseService(opt.Service)
	if err != nil {
		log.Fatal(err)
	}

	_, key := Plugins.SqlcmdFuncMap[s.Schema]
	if !key {
		log.Fatalf("Available Plugins: %s", strings.Join(Plugins.SqlcmdKeys, ","))
	}

	task := model.SqlcmdTask{Ip: s.Ip, Port: s.Port, User: opt.User, Password: opt.Password, SqlcmdPlugin: s.Schema, Query: opt.Query}
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

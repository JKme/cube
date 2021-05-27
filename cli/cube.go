package cli

import (
	"cube/cubelib"
	"cube/log"
	"cube/model"
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

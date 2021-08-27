package cubelib

import (
	"context"
	"cube/log"
	"cube/model"
	Plugins "cube/plugins"
	"cube/util"
	//Plugins "cube/plugins"
	"fmt"
	"strings"
	"sync"
	"time"
)

func validPlugin(plugin string) []string {
	pluginList := strings.Split(plugin, ",")
	if len(pluginList) > 1 && util.SliceContain("ALL", pluginList) {
		log.Errorf("invalid plugin: %s", plugin)
	}
	if plugin == "ALL" {
		pluginList = Plugins.ProbeKeys
	}
	return pluginList
}

func generateTasks(AliveIPS []util.IpAddr, scanPlugin []string) (tasks []model.ProbeTask) {
	tasks = make([]model.ProbeTask, 0)
	for _, aliveAddr := range AliveIPS {
		service := model.ProbeTask{Ip: aliveAddr.Ip, Port: aliveAddr.Port, ScanPlugin: aliveAddr.Plugin}
		tasks = append(tasks, service)
	}

	return tasks
}

func saveReport(taskResult model.ProbeTaskResult) {
	if len(taskResult.Result) > 0 {
		s := fmt.Sprintf("[*]: %s\n[*]: %s:%s\n", taskResult.ProbeTask.ScanPlugin, taskResult.ProbeTask.Ip, taskResult.ProbeTask.Port)
		s1 := fmt.Sprintf("%s\n", taskResult.Result)
		log.Infof(s + s1)
	}
}

func executeProbeTask(ctx context.Context, taskChan chan model.ProbeTask, wg *sync.WaitGroup, delay int) {
	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-taskChan:
			if !ok {
				return
			}

			log.Debugf("Probe %s: %s://%s:%s", task.ScanPlugin, task.ScanPlugin, task.Ip, task.Port)
			fn := Plugins.ProbeFuncMap[task.ScanPlugin]
			r := fn(task)
			saveReport(r)

			select {
			case <-ctx.Done():
			case <-time.After(time.Duration(delay) * time.Second):
			}
			wg.Done()

		}

	}
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

func StartProbeTask(opt *model.ProbeOptions, globalopts *model.GlobalOptions) {
	var (
		threadNum int
		delay     int
	)
	delay = globalopts.Delay
	threadNum = globalopts.Threads
	t1 := time.Now()
	if delay > 0 {
		threadNum = 1
	}

	ips, err := util.ParseIP(opt.Target, opt.TargetFile)
	if err != nil {
		log.Error(err)
	}
	pluginList := validPlugin(opt.ScanPlugin)
	if !util.Subset(pluginList, Plugins.ProbeKeys) && !util.Subset(pluginList, Plugins.ProbeFuncExclude) {
		log.Errorf("plugins not found: %s", pluginList)
	}
	log.Infof("Loading plugin: %s", strings.Join(pluginList, ","))
	ctx := context.Background()

	AliveIPS := util.CheckAlive(ctx, threadNum, delay, ips, pluginList, opt.Port)
	tasks := generateTasks(AliveIPS, pluginList)

	taskChan := make(chan model.ProbeTask, threadNum*2)
	var wg sync.WaitGroup

	//消费者
	for i := 0; i < threadNum; i++ {
		go executeProbeTask(ctx, taskChan, &wg, delay)
	}

	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}
	//wg.Wait()
	waitTimeout(&wg, model.ThreadTimeout)
	util.GetFinishTime(t1)

}

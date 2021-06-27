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

func ValidPlugin(plugin string) ([]string, error) {
	pluginList := strings.Split(plugin, ",")
	if len(pluginList) > 1 && SliceContain("ALL", pluginList) {
		return nil, fmt.Errorf("invalid plugin: %s", plugin)
	}

	if plugin == "ALL" {
		pluginList = Plugins.ProbeKeys
	}
	return pluginList, nil
}

func generateTasks(AliveIPS []util.IpAddr, scanPlugin []string) (tasks []model.ProbeTask) {
	tasks = make([]model.ProbeTask, 0)
	for _, plugin := range scanPlugin {
		for _, aliveAddr := range AliveIPS {
			service := model.ProbeTask{Ip: aliveAddr.Ip, Port: aliveAddr.Port, ScanPlugin: plugin}
			tasks = append(tasks, service)
		}
	}
	return tasks
}

func saveReport(taskResult model.ProbeTaskResult) {
	if len(taskResult.Result) > 0 {
		s := fmt.Sprintf("[*]: %s\n[*]: %s:%s\n", taskResult.ProbeTask.ScanPlugin, taskResult.ProbeTask.Ip, taskResult.ProbeTask.Port)
		s1 := fmt.Sprintf("[*]: %s", taskResult.Result)
		log.Info(s + s1)
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

			log.Debugf("Checking %s Password: %s://%s:%s", task.ScanPlugin, task.ScanPlugin, task.Ip, task.Port)
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
	pluginList, err := ValidPlugin(opt.ScanPlugin)
	if err != nil {
		log.Error(err)
	}
	if !Subset(pluginList, Plugins.ProbeKeys) {
		log.Errorf("plugins not found: %s", pluginList)
	}

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

	waitTimeout(&wg, model.ThreadTimeout)
	getFinishTime(t1)

}

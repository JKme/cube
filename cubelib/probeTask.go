package cubelib

import (
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
		pluginList = Plugins.ProbeKeys[1:]
	}
	return pluginList, nil
}

func GenerateTasks(ipList []string, port int, scanPlugin []string) (tasks []model.ProbeTask) {
	tasks = make([]model.ProbeTask, 0)
	for _, plugin := range scanPlugin {
		for _, ip := range ipList {
			service := model.ProbeTask{Ip: ip, Port: port, ScanPlugin: plugin}
			tasks = append(tasks, service)
		}
	}
	return tasks
}

func saveReport(taskResult model.ProbeTaskResult) {
	if len(taskResult.Result) > 0 {
		s := fmt.Sprintf("[*]: %s\n[*]: %s:%d\n", taskResult.ProbeTask.ScanPlugin, taskResult.ProbeTask.Ip, taskResult.ProbeTask.Port)
		s1 := fmt.Sprintf("[*]: %s", taskResult.Result)
		fmt.Println(s + s1)
	}
}

func executeTask(taskChan chan model.ProbeTask, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskChan {
		//fmt.Println("Hello")
		fn := Plugins.ProbeFuncMap[task.ScanPlugin]
		saveReport(fn(task))
	}

}

func RunTasks(tasks []model.ProbeTask, scanNum int, timeout int) {
	tasksChan := make(chan model.ProbeTask, scanNum*2)
	var wg sync.WaitGroup

	//消费者
	wg.Add(scanNum)
	for i := 0; i < scanNum; i++ {
		go executeTask(tasksChan, &wg)
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
	ips, err := util.ParseIP(opt.Target, opt.TargetFile)
	if err != nil {
		log.Error(err)
	}
	pluginList, err := ValidPlugin(opt.ScanPlugin)
	if err != nil {
		log.Fatal(err)
	}
	if !Subset(pluginList, Plugins.ProbeKeys) {
		log.Fatalf("plugins not found: %s", pluginList)
	}

	tasks := GenerateTasks(ips, opt.Port, pluginList)
	RunTasks(tasks, globalopts.Threads, globalopts.Timeout)
}

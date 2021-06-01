package cubelib

import (
	"cube/model"
	Plugins "cube/plugins"
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
		fmt.Println(strings.Repeat("=", 20))
		fmt.Printf("%s:\n%s", taskResult.ProbeTask.Ip, taskResult.Result)
		fmt.Println(strings.Repeat("=", 20))
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
		//log.Debugf("Put Task Channel: %s", task.Ip)
		tasksChan <- task
	}
	close(tasksChan)

	//wg.Wait()
	waitTimeout(&wg, time.Duration(timeout)*time.Second)
}

//func Scan(scanPlugin string, scanTargets string, scanTargetsFile string, scanPort int, timeout int, scanNum int) {
//	ips, err := ParseIP(scanTargets, scanTargetsFile)
//	if err != nil {
//		log.Error(err)
//	}
//	tasks := generateTasks(ips, scanPort, scanPlugin)
//	runTasks(tasks, scanNum, timeout)
//
//	//aliveIpList := CheckAlive(scanTargets)
//}

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

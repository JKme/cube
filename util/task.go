package util

import (
	"cube/log"
	"cube/model"
	//"fmt"
	"cube/plugins"
	"sync"
	"time"
)



func generateTasks(ipList []string, port int, scanPlugin string)(tasks []model.Task){
	tasks = make([]model.Task, 0)
	for _, ip := range ipList {
		service := model.Task{Ip:ip, Port: port, ScanPlugin: scanPlugin}
		tasks = append(tasks, service)
	}
	return tasks
}

func executeTask(taskChan chan model.Task, wg *sync.WaitGroup){
	defer wg.Done()
	for task :=range taskChan{
		//fmt.Println("Hello")
		fn := Plugins.ScanFuncMap[task.ScanPlugin]
		saveReport(fn(task))
	}

}

func runTasks(tasks []model.Task, scanNum int, timeout int){
	tasksChan := make(chan model.Task, scanNum * 2)
	var wg sync.WaitGroup

	//消费者
	wg.Add(scanNum)
	for i:=0;i<scanNum;i++{
		go executeTask(tasksChan, &wg)
	}

	//生产者
	//go func() {
	//
	//}()

	for _, task := range tasks {
		log.Debugf("Put Task Channel: %s", task.Ip)
		tasksChan <- task
	}
	close(tasksChan)


	//wg.Wait()
	waitTimeout(&wg, time.Duration(timeout) * time.Second)
}

func Scan(scanPlugin string, scanTargets string, scanTargetsFile string, scanPort int, timeout int, scanNum int){
	ips, err := ParseIP(scanTargets, scanTargetsFile)
	if err != nil {
		log.Error(err)
	}
	tasks := generateTasks(ips, scanPort, scanPlugin)
	runTasks(tasks, scanNum, timeout)

	//aliveIpList := CheckAlive(scanTargets)
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
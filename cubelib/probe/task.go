package probe

import (
	"cube/log"
	"cube/model"
	Plugins "cube/plugins"
	"cube/util"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Task struct {
	Ip         string
	Port       int
	ScanPlugin string
}

type TaskResult struct {
	Task   Task
	Result string
	Err    error
}

func generateTasks(ipList []string, port int, scanPlugin string) (tasks []Task) {
	tasks = make([]Task, 0)
	for _, ip := range ipList {
		service := Task{Ip: ip, Port: port, ScanPlugin: scanPlugin}
		tasks = append(tasks, service)
	}
	return tasks
}

func saveReport(taskResult TaskResult) {
	if len(taskResult.Result) > 0 {
		fmt.Println(strings.Repeat("=", 20))
		fmt.Printf("%s:\n%s", taskResult.Task.Ip, taskResult.Result)
		fmt.Println(strings.Repeat("=", 20))
	}
}

func executeTask(taskChan chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskChan {
		//fmt.Println("Hello")
		fn := Plugins.ProbeFuncMap[task.ScanPlugin]
		saveReport(fn(task))
	}

}

func runTasks(tasks []Task, scanNum int, timeout int) {
	tasksChan := make(chan Task, scanNum*2)
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

func Run(opt *model.ProbeOptions, globalopts *model.GlobalOptions) {
	ips, err := util.ParseIP(opt.Target, opt.TargetFile)
	if err != nil {
		log.Error(err)
	}
	tasks := generateTasks(ips, opt.Port, opt.ScanPlugin)
	runTasks(tasks, globalopts.Threads, globalopts.Timeout)
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

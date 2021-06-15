package cubelib

import (
	"context"
	"cube/model"
	Plugins "cube/plugins"
	"fmt"
	"strings"
	"sync"
	"time"
)

//func GenerateCrackTasks(ip []string, port string, auths []model.Auth, plugins []string) (tasks []model.CrackTask) {
//	tasks = make([]model.CrackTask, 0)
//	for _, i := range ip {
//		for _, auth := range auths {
//			for _, p := range plugins {
//				s := model.CrackTask{Ip: i, Port: port, Auth: &auth, CrackPlugin: p}
//				tasks = append(tasks, s)
//			}
//		}
//	}
//	return tasks
//}

func unitTask(ip string, auths []model.Auth, plugins []string) (tasks []model.CrackTask) {
	tasks = make([]model.CrackTask, 0)
	for _, auth := range auths {
		for _, p := range plugins {
			s := model.CrackTask{Ip: ip, Auth: auth, CrackPlugin: p}
			tasks = append(tasks, s)
		}
	}
	return tasks
}

func processArgs(opt *model.CrackOptions) ([]string, error) {

	return nil, nil
}

func generateAuth(user []string, password []string) (authList []model.Auth) {
	authList = make([]model.Auth, 0)
	for _, u := range user {
		for _, pass := range password {
			a := model.Auth{User: u, Password: pass}
			authList = append(authList, a)
		}
	}
	return authList
}

func executeCrackTask(taskChan chan model.CrackTask, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskChan {
		//fmt.Println("Hello")
		fn := Plugins.CrackFuncMap[task.CrackPlugin]
		saveCrackReport(fn(task))
	}

}

func RunCrackTasks(tasks []model.CrackTask, scanNum int, timeout int) {
	tasksChan := make(chan model.CrackTask, scanNum*2)
	var wg sync.WaitGroup

	//消费者
	wg.Add(scanNum)
	for i := 0; i < scanNum; i++ {
		go executeCrackTask(tasksChan, &wg)
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

func runCrackTask(ctx context.Context, taskChan chan model.CrackTask, resultChan chan model.CrackTaskResult, done chan bool) {
	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-taskChan:
			if !ok {
				done <- true
				return
			}
			fn := Plugins.CrackFuncMap[task.CrackPlugin]
			r := fn(task)
			if len(r.Result) > 0 {
				resultChan <- r
				done <- true
			}
		}
	}
}

func monitorResult(done chan bool, resultChan chan model.CrackTaskResult) {
	for r := range resultChan {
		if r.Result != "" {
			fmt.Printf(strings.Repeat("+", 20) + "\n")
			fmt.Printf("Crack Pass Success:%s\n", r)
			fmt.Printf(strings.Repeat("+", 20) + "\n")
			done <- true
		}
	}
}

func saveCrackReport(taskResult model.CrackTaskResult) {
	k := fmt.Sprintf("%v-%v-%v", taskResult.CrackTask.Ip, taskResult.CrackTask.Port, taskResult.CrackTask.CrackPlugin)
	h := MakeTaskHash(k)
	SetTaskHask(h)
	if len(taskResult.Result) > 0 {
		s := fmt.Sprintf("[*]: %s\n[*]: %s:%s\n", taskResult.CrackTask.CrackPlugin, taskResult.CrackTask.Ip, taskResult.CrackTask.Port)
		s1 := fmt.Sprintf("[*]: %s", taskResult.Result)
		fmt.Println(s + s1)
	}
}

func runTask(taskChan chan model.CrackTask, wg *sync.WaitGroup) {
	for task := range taskChan {

		k := fmt.Sprintf("%v-%v-%v", task.Ip, task.Port, task.CrackPlugin)
		h := MakeTaskHash(k)
		if CheckTashHash(h) {
			wg.Done()
			continue
		}

		fn := Plugins.CrackFuncMap[task.CrackPlugin]
		r := fn(task)
		saveCrackReport(r)
		if len(r.Result) > 0 {
			fmt.Println(r)
		}
		wg.Done()
	}
}

func executeIp(ctx context.Context, ip string, authList []model.Auth, plugins []string, resultChan chan model.CrackTaskResult) {

	tasks := unitTask(ip, authList, plugins)
	taskChan := make(chan model.CrackTask, 10)
	//childCtx, childCancel := context.WithCancel(ctx)
	//defer childCancel()
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go func() {
	//
	//	for {
	//		select {
	//		case <-childCtx.Done():
	//			return
	//		//case data, ok := <-resultChan:
	//		//	if ok {
	//		//		fmt.Printf("Get Magic Bean %s\n", data)
	//		//		//fmt.Println(data)
	//		//		childCancel()
	//		//		return
	//		//	}
	//		case task, ok:= <- taskChan:
	//			if ok {
	//for i := 0; i < 2; i++ {

	//for i := 0; i < 3; i++ {
	//	go runTask(taskChan, resultChan)
	//}

	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)
	//runTask(taskChan, resultChan)
}

//func runCrack(plugins []string, ips []string, authList []model.Auth) {
//
//	resultChan := make(chan model.CrackTaskResult, 10)
//	var wg sync.WaitGroup
//
//
//	for _, ip := range ips {
//		Task:
//			for {
//				select {
//				case data, ok := <-resultChan:
//					if ok {
//						fmt.Printf("Get Magic Bean %s\n", data)
//						//fmt.Println(data)
//						break Task
//					}
//				default:
//					tasks := unitTask(ip, authList, plugins)
//					taskChan := make(chan model.CrackTask, 10)
//
//
//					for i:=0;i<3;i++ {
//						wg.Add(1)
//						go runTask(taskChan, resultChan, &wg)
//					}
//
//
//					for _, task := range tasks {
//						fmt.Printf("Put Task: %s\n", task.Auth)
//						taskChan <- task
//					}
//					close(taskChan)
//					wg.Wait()
//				}
//			}
//	}
//
//}

func runCrack(plugins []string, ips []string, authList []model.Auth) {

	for _, ip := range ips {
		var wg sync.WaitGroup
		tasks := unitTask(ip, authList, plugins)
		taskChan := make(chan model.CrackTask, 10)

		go func() {
			for i := 0; i < 4; i++ {
				//go runCrackTask(ctx, taskChan, resultChan,done)
				go runTask(taskChan, &wg)
			}
		}()

		//go func() {
		for _, task := range tasks {
			fmt.Printf("Put Task: %s\n", task.Auth)
			taskChan <- task
		}
		//close(taskChan)
		wg.Wait()
	}

}

// https://stackoverflow.com/questions/45500836/close-multiple-goroutine-if-an-error-occurs-in-one-in-go

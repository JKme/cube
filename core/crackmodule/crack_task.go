package crackmodule

import (
	"context"
	"crypto/md5"
	"cube/config"
	"cube/core"
	"cube/gologger"
	"cube/report"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

func MD5(s string) (m string) {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func MakeTaskHash(k string) string {
	hash := MD5(k)
	return hash
}

func CheckTaskHash(hash string) bool {
	config.SuccessHash.Lock()
	_, ok := config.SuccessHash.S[hash]
	config.SuccessHash.Unlock()
	return ok
}

func SetTaskHash(hash string) {
	config.SuccessHash.Lock()
	config.SuccessHash.S[hash] = true
	config.SuccessHash.Unlock()
}

// ResultMap 当Mysql或者redis空密码的时候，任何密码都正确，会导致密码刷屏

func SetResultMap(r CrackResult) {
	c := fmt.Sprintf("\nCRACK_PLUG: %s\nCRACK_PORT: %s\nCRACK_ADDR: %s\nCRACK_USER: %s\nCRACK_PASS: %s", r.Crack.Name, r.Crack.Port, r.Crack.Ip, r.Crack.Auth.User, r.Crack.Auth.Password)
	data := report.CsvCell{
		Ip:     r.Crack.Ip,
		Module: "Crack_" + r.Crack.Name,
		Cell:   c,
	}
	report.ConcurrentSlices.Append(data)
}

func GetFinishTime(t1 time.Time) {

	fmt.Println(strings.Repeat(">", 50))
	End := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("Finished: %s  Cost: %s", End, time.Since(t1))

}

func WaitThreadTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
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

func buildDefaultTasks(AliveIPS []IpAddr) (cracks []Crack) {
	cracks = make([]Crack, 0)
	for _, addr := range AliveIPS {
		authMaps := GetPluginAuthMap(addr.PluginName)
		auths := authMaps[addr.PluginName]
		for _, auth := range auths {
			s := Crack{Ip: addr.Ip, Port: addr.Port, Auth: auth, Name: addr.PluginName}
			gologger.Debugf("build task: IP:%s  Port:%s  Login:%s  Pass:%s", s.Ip, s.Port, s.Auth.User, s.Auth.Password)
			cracks = append(cracks, s)
		}
	}
	return cracks
}

func buildTasks(AliveIPS []IpAddr, auths []Auth) (cracks []Crack) {
	cracks = make([]Crack, 0)
	for _, addr := range AliveIPS {
		for _, auth := range auths {
			s := Crack{Ip: addr.Ip, Port: addr.Port, Auth: auth, Name: addr.PluginName}
			gologger.Debugf("build task: IP:%s  Port:%s  Login:%s  Pass:%s", s.Ip, s.Port, s.Auth.User, s.Auth.Password)
			cracks = append(cracks, s)
		}
	}
	return cracks
}

func saveCrackResult(crackResult CrackResult) {

	if len(crackResult.Result) > 0 {
		gologger.Debugf("Successful: IP:%s  Port:%s  Login:%s  Pass:%s", crackResult.Crack.Ip, crackResult.Crack.Port, crackResult.Crack.Auth.User, crackResult.Crack.Auth.Password)
		k := fmt.Sprintf("%v-%v-%v", crackResult.Crack.Ip, crackResult.Crack.Port, crackResult.Crack.Name)
		h := MakeTaskHash(k)
		SetTaskHash(h)
		//s1 := fmt.Sprintf("[+]: %s://%s:%s %s", taskResult.CrackTask.CrackPlugin, taskResult.CrackTask.Ip, taskResult.CrackTask.Port, taskResult.Result)
		//fmt.Println(s1)
		SetResultMap(crackResult)

	}

}

func runSingleTask(ctx context.Context, crackTasksChan chan Crack, wg *sync.WaitGroup, delay float64) {
	for {
		select {
		case <-ctx.Done():
			return
		case crackTask, ok := <-crackTasksChan:
			if !ok {
				return
			}
			k := fmt.Sprintf("%v-%v-%v", crackTask.Ip, crackTask.Port, crackTask.Name)
			h := MakeTaskHash(k)
			if CheckTaskHash(h) {
				wg.Done()
				continue
			}
			ic := crackTask.NewICrack()
			gologger.Debugf("cracking: IP:%s  Port:%s  Login:%s  Pass:%s", crackTask.Ip, crackTask.Port, crackTask.Auth.User, crackTask.Auth.Password)
			r := ic.Exec()
			saveCrackResult(r)
			wg.Done()

			select {
			case <-ctx.Done():
			case <-time.After(time.Duration(core.RandomDelay(delay)) * time.Second):
			}
		}
	}
}

func StartCrack(opt *CrackOption, globalopt *core.GlobalOption) {
	var (
		crackPlugins []string
		crackIPS     []string
		crackAuths   []Auth
		crackTasks   []Crack
		threadNum    int
		delay        float64
		aliveIPS     []IpAddr
	)

	ctx := context.Background()
	t1 := time.Now()
	delay = globalopt.Delay
	threadNum = globalopt.Threads

	if delay > 0 {
		//添加使用--delay选项的时候，强制单线程。现在还停留在想象中的攻击
		threadNum = 1
		gologger.Infof("Running in single thread mode when --delay is set")
	}

	crackPlugins = opt.ParsePluginName()
	crackIPS = opt.ParseIP()

	if opt.Port != "" {
		validPort := opt.ParsePort()
		if len(crackPlugins) > 1 && validPort {
			//指定端口的时候仅限定一个插件使用
			gologger.Errorf("plugin is limited to single one when --port is set\n")
		}
	}
	aliveIPS = CheckPort(ctx, threadNum, delay, crackIPS, crackPlugins, opt.Port)

	if len(opt.User+opt.UserFile+opt.Pass+opt.PassFile) > 0 {
		crackAuths = opt.ParseAuth()
		crackTasks = buildTasks(aliveIPS, crackAuths)
	} else {
		crackTasks = buildDefaultTasks(aliveIPS)
	}

	var wg sync.WaitGroup
	taskChan := make(chan Crack, threadNum*2)

	for i := 0; i < threadNum; i++ {
		go runSingleTask(ctx, taskChan, &wg, delay)
	}

	for _, task := range crackTasks {
		wg.Add(1)
		taskChan <- task
	}
	//wg.Wait()

	WaitThreadTimeout(&wg, config.ThreadTimeout)
	for k := range report.ConcurrentSlices.Iter() {
		gologger.Infof("%s", k.Value.Cell)

	}

	GetFinishTime(t1)
	srMap := make(map[string]int)
	for k := range report.ConcurrentSlices.Iter() {
		sr := k.Value.Module
		srMap[sr] += 1
		fmt.Println(srMap)
	}
}

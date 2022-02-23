package plugins

import (
	"context"
	"crypto/md5"
	"cube/conf"
	"cube/core"
	"cube/gologger"
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
	conf.SuccessHash.Lock()
	_, ok := conf.SuccessHash.S[hash]
	conf.SuccessHash.Unlock()
	//log.Debugf("Success: %#v\n", model.SuccessHash)
	return ok
}

func SetTaskHash(hash string) {
	conf.SuccessHash.Lock()
	conf.SuccessHash.S[hash] = true
	conf.SuccessHash.Unlock()
}

// ResultMap 当Mysql或者redis空密码的时候，任何密码都正确，会导致密码刷屏
var ResultMap = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

func SetResultMap(r CrackResult) {
	ResultMap.Lock()
	ResultMap.m[fmt.Sprintf("%s==>%s:%s", r.Crack.Name, r.Crack.Ip, r.Crack.Port)] = r.Result
	ResultMap.Unlock()
}

func ReadResultMap() {
	ResultMap.RLock()
	n := ResultMap.m
	ResultMap.RUnlock()
	for k, v := range n {
		gologger.Infof("[*]: %s %v", k, v)
	}
}

func GetFinishTime(t1 time.Time) {

	fmt.Println(strings.Repeat(">", 50))
	End := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("Finished:%s  Cost:%s", End, time.Since(t1))

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
			cracks = append(cracks, s)
		}
	}
	return cracks
}

func saveResult() {

}

func saveCrackReport(crackResult CrackResult) {

	if len(taskResult.Result) > 0 {
		gologger.Debugf("Put Result to Map: %v\n", taskResult)
		k := fmt.Sprintf("%v-%v-%v", taskResult.CrackTask.Ip, taskResult.CrackTask.Port, taskResult.CrackTask.CrackPlugin)
		h := util.MakeTaskHash(k)
		util.SetTaskHash(h)
		//s1 := fmt.Sprintf("[+]: %s://%s:%s %s", taskResult.CrackTask.CrackPlugin, taskResult.CrackTask.Ip, taskResult.CrackTask.Port, taskResult.Result)
		//fmt.Println(s1)
		util.SetResultMap(taskResult)
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
			//if task.Port == "" {
			//	task.Port =  strconv.Itoa(model.CommonPortMap[task.CrackPlugin])
			//}
			//alive := CheckAlive(task)
			//if !alive {
			//	wg.Done()
			//	continue
			//}

			//gologger.Debugf("Cracking %s password ", crackTask.Ip, crackTask.CrackPlugin, task.Auth.User, task.Auth.Password, task.Ip, task.Port)
			k := fmt.Sprintf("%v-%v-%v", crackTask.Ip, crackTask.Port, crackTask.Name)
			h := MakeTaskHash(k)
			if CheckTaskHash(h) {
				wg.Done()
				continue
			}
			c := crackTask.NewICrack()
			r := c.Exec()
			//fn := CrackFuncMap[task.CrackPlugin]
			//r := fn(task)
			saveCrackReport(r)
			wg.Done()

			select {
			case <-ctx.Done():
			case <-time.After(time.Duration(delay) * time.Second):
			}

		}

	}
}

func parseAuthOption() {

}

func parsePluginOption() {

}

func getAuthList() {

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

	if delay > 0 {
		//添加使用--delay选项的时候，强制单线程。现在还停留在想象中的攻击
		threadNum = 1
	} else {
		threadNum = globalopt.Threads
	}

	crackPlugins = opt.ParsePluginName()
	crackIPS = opt.ParseIP()

	if opt.Port != "" {
		validPort := opt.ParsePort()
		if len(crackPlugins) > 1 && validPort {
			//指定端口的时候仅限定一个插件使用
			gologger.Errorf("plugins are limited to single one when --port is set\n")
		}
	}
	aliveIPS = CheckPort(ctx, threadNum, delay, crackIPS, crackPlugins, opt.Port)

	gologger.Infof("crackPlugins: %s\n", crackPlugins)
	gologger.Infof("crackIPS: %s\n", crackIPS)

	if len(opt.User+opt.UserFile+opt.Pass+opt.PassFile) > 0 {
		crackAuths = opt.ParseAuth()
		crackTasks = buildTasks(aliveIPS, crackAuths)
	} else {
		crackTasks = buildDefaultTasks(aliveIPS)
	}
	gologger.Debugf("build tasks: %v", crackTasks)
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
	WaitThreadTimeout(&wg, conf.ThreadTimeout*2)
	ReadResultMap()
	GetFinishTime(t1)
	//if util.Contains(opt.CrackPluginName, Plugins.CrackFuncExclude) {
	//	//当-x是单独使用的插件，比如phpmyadmin、basicAuth类型的时候
	//	AliveIPS = append(AliveIPS, util.IpAddr{
	//		Ip:     opt.Ip,
	//		Port:   "",
	//		Plugin: opt.CrackPlugin,
	//	})
	//} else {
	//	optPlugins = genPlugins(opt.CrackPlugin)
	//	gologger.Infof("Loading plugin: %s", strings.Join(optPlugins, ","))
	//	ips, _ = util.ParseIP(opt.Ip, opt.IpFile)
	//
	//	AliveIPS = util.CheckAlive(ctx, num, delay, ips, optPlugins, opt.Port)
	//	//AliveIPS = RemoveRepByMap(AliveIPS)  // 去重IP
	//	gologger.Debugf("Receive alive IP: %s", AliveIPS)
	//}

	gologger.Debugf(string(rune(threadNum)), aliveIPS, ctx, t1, crackPlugins, crackIPS, crackAuths)
}

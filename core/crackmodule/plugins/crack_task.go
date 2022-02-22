package plugins

import (
	"context"
	"crypto/md5"
	"cube/conf"
	"cube/core"
	"cube/gologger"
	"cube/pkg/util"
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

func buildTasks() {
	// 生成任务
}

func saveResult() {

}

func runSingleTask() {

}

func parseAuthOption() {

}

func parsePluginOption() {

}

func getAuthList() {

}

func StartCrack(opt *CrackOptions, globalopts *core.GlobalOptions) {
	var (
		//optPlugins []string
		//ips        []string
		//auths      []Auth
		//crackTasks []Crack
		num      int
		delay    float64
		aliveIPS []util.IpAddr
	)
	ctx := context.Background()
	t1 := time.Now()
	delay = globalopts.Delay

	if delay > 0 {
		//添加使用--delay选项的时候，强制单线程。PS：用到的时候提个ISSUE，现在还停留在想象中的攻击
		num = 1
	} else {
		num = globalopts.Threads
	}

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

	gologger.Debugf(opt.CrackPluginName)
	gologger.Debugf(string(rune(num)), aliveIPS, ctx, t1)
}

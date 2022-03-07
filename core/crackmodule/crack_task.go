package crackmodule

import (
	"bytes"
	"context"
	"crypto/md5"
	"cube/config"
	"cube/core"
	"cube/gologger"
	"cube/pkg"
	"cube/report"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var SuccessHash = struct {
	sync.RWMutex
	S map[string]bool
}{S: make(map[string]bool)}

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
	SuccessHash.Lock()
	_, ok := SuccessHash.S[hash]
	SuccessHash.Unlock()
	return ok
}

func SetTaskHash(hash string) {
	SuccessHash.Lock()
	SuccessHash.S[hash] = true
	SuccessHash.Unlock()
}

func CrackHelpTable() string {
	buf := bytes.NewBufferString("")
	var flag = ""
	table := tablewriter.NewWriter(buf)
	table.SetHeader([]string{"Func", "Port", "Load By X"})
	for _, k := range CrackKeys {
		if pkg.Contains(k, config.CrackX) {
			flag = "Y"
		} else {
			flag = "N"
		}
		table.Append([]string{k, GetCrackPort(k), flag})
		table.SetRowLine(true)
	}
	table.Render()
	return buf.String()
}

// ResultMap 当Mysql或者redis空密码的时候，任何密码都正确，会导致密码刷屏

func SetResultMap(r CrackResult) {
	var c string
	if len(r.Extra) > 0 {
		c = fmt.Sprintf("\nCRACK_PLUG: %s\nCRACK_PORT: %s\nCRACK_ADDR: %s\nCRACK_USER: %s\nCRACK_PASS: %s\nCRACKEXTRA: %s", r.Crack.Name, r.Crack.Port, r.Crack.Ip, r.Crack.Auth.User, r.Crack.Auth.Password, r.Extra)
	} else {
		c = fmt.Sprintf("\nCRACK_PLUG: %s\nCRACK_PORT: %s\nCRACK_ADDR: %s\nCRACK_USER: %s\nCRACK_PASS: %s", r.Crack.Name, r.Crack.Port, r.Crack.Ip, r.Crack.Auth.User, r.Crack.Auth.Password)
	}

	data := report.CsvCell{
		Ip:     r.Crack.Ip,
		Module: "Crack_" + r.Crack.Name,
		Cell:   c,
	}
	report.ConcurrentSlices.Append(data)
	report.CsvShells = append(report.CsvShells, data)
}

func GetFinishTime(t1 time.Time) {

	fmt.Println(strings.Repeat(">", 50))
	End := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("Finished: %s  Cost: %s\n", End, time.Since(t1))

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
		gologger.Debug("threading timeout, target service is not available")
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
			gologger.Debugf("cracking %s: IP:%s  Port:%s  Login:%s  Pass:%s", crackTask.Name, crackTask.Ip, crackTask.Port, crackTask.Auth.User, crackTask.Auth.Password)
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
		fp           string
	)

	ctx := context.Background()
	t1 := time.Now()
	delay = globalopt.Delay
	threadNum = globalopt.Threads
	fp = globalopt.Output

	if delay > 0 {
		//添加使用--delay选项的时候，强制单线程。现在还停留在想象中的攻击
		threadNum = 1
		gologger.Infof("Running in single thread mode when --delay is set")
	}

	crackPlugins = opt.ParsePluginName()
	if len(crackPlugins) == 0 {
		gologger.Errorf("plug doesn't exist: %s", opt.PluginName)
	}
	gologger.Debugf("load plug: %s", crackPlugins)
	if len(crackPlugins) == 1 && GetMutexStatus(crackPlugins[0]) {
		// phpmyadmin、httpbasic之类的跳过IP检查
		crackIPS = []string{opt.Ip}
		aliveIPS = append(aliveIPS, IpAddr{
			Ip:         crackIPS[0],
			Port:       "",
			PluginName: crackPlugins[0],
		})
	} else {
		crackIPS = opt.ParseIP()
		if opt.Port != "" {
			validPort := opt.ParsePort()
			if len(crackPlugins) > 1 && validPort {
				//指定端口的时候仅限定一个插件使用
				gologger.Errorf("plugin is limited to single one when --port is set\n")
			}
		}
		aliveIPS = CheckPort(ctx, threadNum, delay, crackIPS, crackPlugins, opt.Port)
	}
	if len(crackIPS) == 0 {
		gologger.Errorf("target service is missing, please set -s/-S")
	}

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
	ccs := report.RemoveDuplicateCSS(report.CsvShells)
	r := report.RemoveDuplicateResult(ccs)
	for _, v := range r {
		gologger.Infof("%s", v.Cell)
	}

	if len(fp) > 0 {
		if _, err := os.Stat(fp); err == nil {
			// path/to/whatever exists
			cs := report.ReadExportExcel(fp)
			gologger.Infof("Appending result to %s success", fp)
			for _, v := range cs {
				report.CsvShells = append(report.CsvShells, v)
				//gologger.Debugf("Appending %s", v.Ip)
			}
			css2 := report.RemoveDuplicateCSS(report.CsvShells)
			report.WriteExportExcel(css2, fp)
		} else if errors.Is(err, os.ErrNotExist) {
			// path/to/whatever does *not* exist
			report.WriteExportExcel(report.CsvShells, fp)
			gologger.Infof("Write result to %s success", fp)

		} else {
			// Schrodinger: file may or may not exist. See err for details.

			// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
			gologger.Errorf("can't find file %s, err: %s", fp, err)
		}
	}
	GetFinishTime(t1)
}

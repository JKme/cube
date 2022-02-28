package probemodule

import (
	"context"
	"cube/config"
	"cube/core"
	"cube/core/crackmodule"
	"cube/gologger"
	"cube/report"
	"fmt"
	"sync"
	"time"
)

func buildTasks(aliveIPS []IpAddr, scanPlugins []string) (probes []Probe) {
	probes = make([]Probe, 0)
	for _, addr := range aliveIPS {
		service := Probe{Ip: addr.Ip, Port: addr.Port, Name: addr.PluginName}
		probes = append(probes, service)
	}
	return probes
}

func SetResult(probeResult ProbeResult) {
	if len(probeResult.Result) > 0 {
		c := fmt.Sprintf("[*]: %s\n[*]: %s:%s\n%s\n", probeResult.Probe.Name, probeResult.Probe.Ip, probeResult.Probe.Port, probeResult.Result)

		data := report.CsvCell{
			Ip:     probeResult.Probe.Ip,
			Module: "Probe_" + probeResult.Probe.Name,
			Cell:   c,
		}
		report.ConcurrentSlices.Append(data)
		report.CsvShells = append(report.CsvShells, data)
	}
}

func runSingleTask(ctx context.Context, taskChan chan Probe, wg *sync.WaitGroup, delay float64) {
	for {
		select {
		case <-ctx.Done():
			return
		case probeTask, ok := <-taskChan:
			if !ok {
				return
			}
			ic := probeTask.NewIProbe()
			gologger.Debugf("probing: IP:%s  Port:%s", probeTask.Ip, probeTask.Port)
			r := ic.ProbeExec()
			SetResult(r)

			select {
			case <-ctx.Done():
			case <-time.After(time.Duration(core.RandomDelay(delay)) * time.Second):
			}
			wg.Done()
		}
	}
}

func StartProbe(opt *ProbeOption, globalopt *core.GlobalOption) {
	var (
		probePlugins []string
		probeIPS     []string
		probeTasks   []Probe
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

	probePlugins = opt.ParsePluginName()
	probeIPS = opt.ParseIP()
	if opt.Port != "" {
		validPort := opt.ParsePort()
		if len(probePlugins) > 1 && validPort {
			//指定端口的时候仅限定一个插件使用
			gologger.Errorf("plugin is limited to single one when --port is set\n")
		}
	}

	aliveIPS = CheckPort(ctx, threadNum, delay, probeIPS, probePlugins, opt.Port)
	probeTasks = buildTasks(aliveIPS, probePlugins)

	var wg sync.WaitGroup
	taskChan := make(chan Probe, threadNum*2)

	//消费者
	for i := 0; i < threadNum; i++ {
		go runSingleTask(ctx, taskChan, &wg, delay)
	}

	for _, task := range probeTasks {
		wg.Add(1)
		taskChan <- task
	}
	//wg.Wait()
	crackmodule.WaitThreadTimeout(&wg, config.ThreadTimeout)

	for k := range report.ConcurrentSlices.Iter() {
		gologger.Infof("%s", k.Value.Cell)

	}

	plugMap := report.SortPlug()
	ipMap := report.SortIP()

	plugsElement := report.GetKeys(plugMap)
	plugsElement = append([]string{"IP"}, plugsElement...)

	crackmodule.GetFinishTime(t1)
	report.ExportTitle(plugsElement, report.GetKeys(ipMap), report.CsvShells)
}

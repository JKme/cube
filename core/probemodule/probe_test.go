package probemodule

import (
	"fmt"
	"github.com/lcvvvv/gonmap"
	"testing"
	"time"
)

func TestGoNmap(t *testing.T) {
	gonmap.Init(9)
	gonmap.SetTimeout(time.Duration(3) * time.Second)
	//开启全探针模式
	gonmap.SetScanVersion()
	var addr = "172.20.40.1"
	var port = 22
	//
	tcpBanner := gonmap.GetTcpBanner(addr, port, gonmap.New(), time.Duration(3)*time.Second)
	fmt.Println(tcpBanner)
	//port1 := tcpBanner.Target.Port()
	//addr1 := tcpBanner.Target.Addr()
	finger := tcpBanner.TcpFinger
	r := gonmap.GetAppBannerFromTcpBanner(tcpBanner)
	fmt.Println(r)
	fmt.Println(finger)
}

func TestPrometheus_ProbeExec(t *testing.T) {
	p := Probe{
		Ip:   "172.21.96.220",
		Port: "9090",
		Name: "prometheus",
	}

	task := p.NewIProbe()
	task.ProbeExec()
}

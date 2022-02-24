package crackmodule

import (
	"context"
	"cube/config"
	"cube/core"
	"cube/gologger"
	"fmt"
	"net"
	"sync"
	"time"
)

type IpAddr struct {
	Ip         string
	Port       string
	PluginName string
}

var (
	mutex     sync.Mutex
	AliveAddr []IpAddr
	ipList    []IpAddr
)

func CheckPort(ctx context.Context, threadNum int, delay float64, ips []string, pluginNames []string, port string) []IpAddr {
	//指定插件端口的时候，只允许加载一个插件
	if len(port) > 0 {
		for _, ip := range ips {
			ipList = append(ipList, IpAddr{
				Ip:         ip,
				Port:       port,
				PluginName: pluginNames[0],
			})
		}
	} else {
		for _, plugin := range pluginNames {
			for _, ip := range ips {
				ipList = append(ipList, IpAddr{
					Ip:         ip,
					Port:       GetPort(plugin),
					PluginName: plugin,
				})
			}
		}

	}

	var addrChan = make(chan IpAddr, threadNum*2)
	var wg sync.WaitGroup
	wg.Add(len(ipList))

	for i := 0; i < threadNum; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case addr, ok := <-addrChan:
					if !ok {
						return
					}
					if GetTCP(addr.PluginName) {
						//TCP的时候检查端口，UDP跳过
						SaveAddr(check(addr))
					}
					wg.Done()
					select {
					case <-ctx.Done():
					case <-time.After(time.Duration(core.RandomDelay(delay)) * time.Second):
					}
				}
			}
		}()
	}

	for _, addr := range ipList {
		addrChan <- addr
	}
	close(addrChan)
	wg.Wait()

	return AliveAddr
}

func check(addr IpAddr) (bool, IpAddr) {
	alive := false
	gologger.Debugf("tcp port conn check: %s:%s", addr.Ip, addr.Port)
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", addr.Ip, addr.Port), config.TcpConnTimeout)
	if err == nil {
		gologger.Infof("Open %s:%s", addr.Ip, addr.Port)
		alive = true
	}
	return alive, addr
}

//func checkUDP(addr IpAddr) (bool, IpAddr) {
//	//https://github.com/bronzdoc/gops
//	//alive := true
//	gologger.Debugf("skip udp port conn check: %s:%s", addr.Ip, addr.Port)
//	time.Sleep(time.Millisecond * 10)
//
//	return true, addr
//}

func SaveAddr(alive bool, addr IpAddr) {
	if alive {
		mutex.Lock()
		AliveAddr = append(AliveAddr, addr)
		mutex.Unlock()
	}
}

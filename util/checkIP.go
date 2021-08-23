package util

import (
	"context"
	"cube/log"
	"cube/model"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type IpAddr struct {
	Ip     string
	Port   string
	Plugin string
}

var (
	mutex     sync.Mutex
	AliveAddr []IpAddr
	ipList    []IpAddr
)

func CheckAlive(ctx context.Context, Num int, delay int, ips []string, plugins []string, port string) []IpAddr {
	if len(port) > 0 {
		for _, ip := range ips {
			ipList = append(ipList, IpAddr{
				Ip:   ip,
				Port: port,
			})
		}
	} else {
		for _, plugin := range plugins {
			for _, ip := range ips {
				ipList = append(ipList, IpAddr{
					Ip:     ip,
					Port:   strconv.Itoa(model.CommonPortMap[plugin]),
					Plugin: plugin,
				})
			}
		}

	}

	//var wg sync.WaitGroup
	//wg.Add(len(ipList))
	//
	//for _, addr := range ipList {
	//	go func(addr IpAddr) {
	//		defer wg.Done()
	//		SaveAddr(check(addr))
	//	}(addr)
	//}
	//wg.Wait()
	var threadNum int
	if delay != 0 {
		threadNum = 1
	} else {
		threadNum = Num
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
					if addr.Plugin == "netbios" {
						SaveAddr(checkUDP(addr))
					} else {
						SaveAddr(check(addr))
					}
					wg.Done()
					select {
					case <-ctx.Done():
					case <-time.After(time.Duration(delay) * time.Second):
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
	log.Debugf("Port connect check: %s:%s", addr.Ip, addr.Port)
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", addr.Ip, addr.Port), model.ConnectTimeout)
	if err == nil {
		log.Infof("%s:%s Open", addr.Ip, addr.Port)
		alive = true
	}
	return alive, addr
}

func checkUDP(addr IpAddr) (bool, IpAddr) {
	alive := false
	log.Debugf("Port connect check: %s:%s", addr.Ip, addr.Port)
	_, err := net.DialTimeout("udp", fmt.Sprintf("%v:%v", addr.Ip, addr.Port), model.ConnectTimeout)
	if err == nil {
		log.Infof("%s:%s Open", addr.Ip, addr.Port)
		alive = true
	}
	return alive, addr
}

func SaveAddr(alive bool, addr IpAddr) {
	if alive {
		mutex.Lock()
		AliveAddr = append(AliveAddr, addr)
		mutex.Unlock()
	}
}

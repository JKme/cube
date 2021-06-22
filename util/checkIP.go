package util

import (
	"cube/log"
	"cube/model"
	"fmt"
	"net"
	"strconv"
	"sync"
)

type IpAddr struct {
	Ip   string
	Port string
}

var (
	mutex     sync.Mutex
	AliveAddr []string
	ipList    []IpAddr
)

func CheckAlive(ips []string, plugins []string, port string) []string {
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
					Ip:   ip,
					Port: strconv.Itoa(model.CommonPortMap[plugin]),
				})
			}
		}

	}

	fmt.Println(ipList)
	var wg sync.WaitGroup
	wg.Add(len(ipList))

	for _, addr := range ipList {
		go func(addr IpAddr) {
			defer wg.Done()
			SaveAddr(check(addr))
		}(addr)
	}
	wg.Wait()

	return AliveAddr
}

func check(addr IpAddr) (bool, IpAddr) {
	alive := false
	log.Debugf("Checking: %s:%s", addr.Ip, addr.Port)
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", addr.Ip, addr.Port), model.T)
	if err == nil {
		alive = true
	}
	return alive, addr
}

func SaveAddr(alive bool, addr IpAddr) {
	if alive {
		mutex.Lock()
		AliveAddr = append(AliveAddr, addr.Ip)
		mutex.Unlock()
	}
}

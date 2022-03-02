package probemodule

import (
	"cube/config"
	"cube/core/crackmodule"
	"cube/gologger"
	"cube/pkg"
	"strconv"
	"strings"
)

type ProbeOption struct {
	Ip         string
	IpFile     string
	Port       string
	PluginName string
}

func NewProbeOption() *ProbeOption {
	return &ProbeOption{}
}

func (po *ProbeOption) ParsePluginName() []string {
	var pluginNameList []string

	pns := strings.Split(po.PluginName, ",")
	if len(pns) > 2 && pkg.Contains("X", pns) {
		//指定-X只能单独加载
		pluginNameList = nil
	}

	if len(pns) == 1 {
		if pns[0] == "X" {
			pluginNameList = config.ProbeX
		}
		if pkg.Contains(pns[0], ProbeKeys) {
			pluginNameList = pns
		}
	} else {
		for _, k := range pns {
			if pkg.Contains(k, ProbeKeys) {
				pluginNameList = append(pluginNameList, k)
			}
		}
	}
	return pluginNameList
}

func (po *ProbeOption) ParseIP() []string {
	var hosts []string
	ip := po.Ip
	fp := po.IpFile

	if ip != "" {
		hosts = crackmodule.ExpandIp(ip)
	}

	if fp != "" {
		var ips []string
		ips, _ = crackmodule.ReadIPFile(fp)
		hosts = append(hosts, ips...)
	}
	hosts = pkg.RemoveDuplicate(hosts)
	return hosts
}

func (po *ProbeOption) ParsePort() bool {
	b, err := strconv.ParseBool(po.Port)
	if err != nil {
		gologger.Errorf("error while parse port option: %v", po.Port)
	}
	return b
}

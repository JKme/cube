package probemodule

import (
	"cube/core/crackmodule"
	"cube/gologger"
	"cube/pkg/util"
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
	if len(pns) > 2 && util.Contains("X", pns) {
		//指定-X只能单独加载
		pluginNameList = nil
	}

	if len(pns) == 1 {
		if pns[0] == "X" {
			for _, k := range ProbeKeys {
				if GetLoadStatus(k) {
					pluginNameList = append(pluginNameList, k)
				}
			}
		}
		if util.Contains(pns[0], ProbeKeys) {
			pluginNameList = pns
		}
	} else {
		for _, k := range pns {
			if util.Contains(k, ProbeKeys) {
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
	hosts = util.RemoveDuplicate(hosts)
	return hosts
}

func (po *ProbeOption) ParsePort() bool {
	b, err := strconv.ParseBool(po.Port)
	if err != nil {
		gologger.Errorf("error while parse port option: %v", po.Port)
	}
	return b
}

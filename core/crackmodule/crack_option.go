package crackmodule

import (
	"bufio"
	"cube/gologger"
	"cube/pkg/util"
	"github.com/malfunkt/iprange"
	"os"
	"strconv"
	"strings"
)

type CrackOption struct {
	Ip         string
	IpFile     string
	User       string
	UserFile   string
	Pass       string
	PassFile   string
	Port       string
	PluginName string
}

func NewCrackOptions() *CrackOption {
	return &CrackOption{}
}

func (cp *CrackOption) ParsePluginName() []string {
	var pluginNameList []string

	pns := strings.Split(cp.PluginName, ",")
	if len(pns) > 2 && util.Contains("X", pns) {
		//指定-X只能单独加载
		pluginNameList = nil
	}
	if len(pns) > 2 && util.Contains("Y", pns) {
		pluginNameList = nil
	}
	switch {
	case len(pns) == 1:
		if pns[0] == "X" {
			for _, k := range CrackKeys {
				if !GetMutexStatus(k) && GetLoadStatus(k) {
					pluginNameList = append(pluginNameList, k)
				}
			}
		}
		if pns[0] == "Y" {
			for _, k := range CrackKeys {
				if !GetMutexStatus(k) {
					pluginNameList = append(pluginNameList, k)
				}
			}
		}
		if util.Contains(pns[0], CrackKeys) {
			pluginNameList = pns
		}
	default:
		for _, k := range pns {
			if util.Contains(k, CrackKeys) {
				pluginNameList = append(pluginNameList, k)
			}
		}
	}
	return pluginNameList
}

func (cp *CrackOption) ParseAuth() []Auth {
	var auths []Auth
	user := cp.User
	userFile := cp.UserFile
	pass := cp.Pass
	passFile := cp.PassFile
	us := opt2slice(user, userFile)
	ps := opt2slice(pass, passFile)

	for _, u := range us {
		for _, p := range ps {
			auths = append(auths, Auth{
				User:     u,
				Password: p,
			})
		}
	}
	return auths
}

func (cp *CrackOption) ParseIP() []string {
	var hosts []string
	ip := cp.Ip
	fp := cp.IpFile

	if ip != "" {
		hosts = ExpandIp(ip)
	}

	if fp != "" {
		var ips []string
		ips, _ = ReadIPFile(fp)
		hosts = append(hosts, ips...)
	}
	hosts = util.RemoveDuplicate(hosts)
	return hosts
}

func (cp *CrackOption) ParsePort() bool {
	b, err := strconv.ParseBool(cp.Port)
	if err != nil {
		gologger.Errorf("error while parse port option: %v", cp.Port)
	}
	return b
}

func opt2slice(str, file string) []string {
	if len(str+file) == 0 {
		gologger.Errorf("Provide login name(-l/-L) and login password(-p/-P)")
	}
	if len(str) > 0 {
		r := strings.Split(str, ",")

		return r
	}
	r := util.FileReader(file)
	return r
}

func ExpandIp(ip string) (hosts []string) {

	list, err := iprange.ParseList(ip)
	if err != nil {
		gologger.Errorf("IP parsing error\nformat: 10.0.0.1, 10.0.0.5-10, 192.168.1.*, 192.168.10.0/24")
	}
	rng := list.Expand()
	for _, v := range rng {
		hosts = append(hosts, v.String())

	}
	return hosts

}

func ReadIPFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		gologger.Debugf("Open %s error, %s\n", filename, err)
	}
	defer file.Close()
	var content []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			host := ExpandIp(text)
			content = append(content, host...)
		}
	}
	return content, nil
}

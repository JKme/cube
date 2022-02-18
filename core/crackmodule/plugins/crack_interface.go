package plugins

import (
	"cube/gologger"
	"strings"
)

var CrackKeys []string

type ICrack interface {
	SetName() string       // 插件名称
	SetPort() string       //设置端口
	SetAuthUser() []string //设置默认爆破的用户名
	SetAuthPass() []string //设置默认爆破的密码
	IsLoad() bool          // ALL选项的时候，是否加载
	Exec() CrackResult     //运行task，返回string
}

func AddCrackKeys(s string) {
	CrackKeys = append(CrackKeys, s)
}

//type Result interface {
//	ResultToString() (string, error)  //probe、crack、sqlcmd都实现获取结果的接口
//}
//

//var ICrackMap map[string]ICrack
//
//func init() {
//	ICrackMap = make(map[string]ICrack)
//}

func (c *Crack) NewICrack() ICrack {
	switch c.Name {
	case "ssh":
		return &SshCrack{c}
	case "ftp":
		return &FtpCrack{c}
	default:
		return nil
	}
}

func NewCrack(s string) Crack {
	return Crack{
		Name: s,
	}
}

func GetPort(s string) string {
	c := NewCrack(s)
	ic := c.NewICrack()
	return ic.SetPort()
}

func GetLoadStatus(s string) string {
	c := NewCrack(s)
	ic := c.NewICrack()
	if ic.IsLoad() == true {
		return "Y"
	} else {
		return "N"
	}
}

func getPluginAuthUser(s string) []string {
	c := NewCrack(s)
	ic := c.NewICrack()
	return ic.SetAuthUser()
}

func getPluginAuthPass(s string) []string {
	c := NewCrack(s)
	ic := c.NewICrack()
	return ic.SetAuthPass()
}

func GetPluginAuths(p string) []Auth {
	auths := make([]Auth, 0)
	for _, user := range getPluginAuthUser(p) {
		for _, pass := range getPluginAuthPass(p) {
			gologger.Debugf("% is preparing credentials: %s <=> %s", p, user, pass)
			pass = strings.Replace(pass, "{user}", user, -1)
			auths = append(auths, Auth{
				User:     user,
				Password: pass,
			})
		}
	}
	return auths
}

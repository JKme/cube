package plugins

import (
	"cube/gologger"
	"strings"
)

type Auth struct {
	User     string
	Password string
}
type Crack struct {
	Ip   string
	Port string
	Auth Auth
	Name string
}

type CrackResult struct {
	Crack  Crack
	Result string
	Err    error
}

var CrackKeys []string

type ICrack interface {
	SetName() string       //插件名称
	SetPort() string       //设置端口
	SetAuthUser() []string //设置默认爆破的用户名
	SetAuthPass() []string //设置默认爆破的密码
	IsLoad() bool          //-x X选项的时候，是否加载
	IsMutex() bool         //只能单独使用的插件，比如phpmyadmin
	IsTcp() bool           //TCP需要先进行端口开放探测，UDP跳过
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

func GetMutexStatus(s string) bool {
	c := NewCrack(s)
	ic := c.NewICrack()
	return ic.IsMutex()
}

func GetTCP(s string) bool {
	c := NewCrack(s)
	ic := c.NewICrack()
	return ic.IsTcp()
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

func GetPluginAuthMap(s string) map[string][]Auth {
	auths := make([]Auth, 0)
	authMaps := make(map[string][]Auth, 0)
	for _, user := range getPluginAuthUser(s) {
		for _, pass := range getPluginAuthPass(s) {
			gologger.Debugf("%s is preparing credentials: %s <=> %s", s, user, pass)
			pass = strings.Replace(pass, "{user}", user, -1)
			auths = append(auths, Auth{
				User:     user,
				Password: pass,
			})
		}
	}
	authMaps[s] = auths
	return authMaps
}

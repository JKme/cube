package crackmodule

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

func GetLoadStatus(s string) bool {
	c := NewCrack(s)
	ic := c.NewICrack()
	return ic.IsLoad()
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

func getPluginAuthCred(s string) bool {
	//检查插件是否设置了默认的用户和密码
	if len(getPluginAuthPass(s)) == 0 || len(getPluginAuthPass(s)) == 0 {
		return false
	}
	return true
}

func GetPluginAuthMap(s string) map[string][]Auth {
	auths := make([]Auth, 0)
	authMaps := make(map[string][]Auth, 0)
	credStatus := getPluginAuthCred(s)
	if !credStatus {
		gologger.Errorf("SetAuthUser() or SetAuthPass() is Empty for %s", s)
	}
	for _, user := range getPluginAuthUser(s) {
		for _, pass := range getPluginAuthPass(s) {
			pass = strings.Replace(pass, "{user}", user, -1)
			gologger.Debugf("%s is preparing default credentials: %s <=> %s", s, user, pass)
			auths = append(auths, Auth{
				User:     user,
				Password: pass,
			})
		}
	}
	authMaps[s] = auths
	return authMaps
}
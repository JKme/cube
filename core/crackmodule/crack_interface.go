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
	Extra  string
	Err    error
}

var CrackKeys []string

type ICrack interface {
	CrackName() string       //插件名称
	CrackPort() string       //设置端口
	CrackAuthUser() []string //设置默认爆破的用户名
	CrackAuthPass() []string //设置默认爆破的密码
	IsMutex() bool           //只能单独使用的插件，比如phpmyadmin
	CrackPortCheck() bool    //TCP需要先进行端口开放探测，UDP跳过
	Exec() CrackResult       //运行task，返回string
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
	case "redis":
		return &Redis{c}
	case "elastic":
		return &Elastic{c}
	case "httpbasic":
		return &FtpCrack{c}
	case "jenkins":
		return &FtpCrack{c}
	case "mongo":
		return &FtpCrack{c}
	case "mssql":
		return &FtpCrack{c}
	case "mysql":
		return &FtpCrack{c}
	case "postgres":
		return &FtpCrack{c}
	case "smb":
		return &FtpCrack{c}
	case "zabbix":
		return &FtpCrack{c}
	case "phpmyadmin":
		return &Phpmyadmin{c}
	case "oracle":
		return &Oracle{c}
	default:
		return nil
	}
}

func NewCrack(s string) Crack {
	return Crack{
		Name: s,
	}
}

func GetCrackPort(s string) string {
	c := NewCrack(s)
	ic := c.NewICrack()
	return ic.CrackPort()
}

func GetMutexStatus(s string) bool {
	c := NewCrack(s)
	ic := c.NewICrack()
	return ic.IsMutex()
}

func NeedPortCheck(s string) bool {
	c := NewCrack(s)
	ic := c.NewICrack()
	return ic.CrackPortCheck()
}

func getPluginAuthUser(s string) []string {
	c := NewCrack(s)
	ic := c.NewICrack()
	return ic.CrackAuthUser()
}

func getPluginAuthPass(s string) []string {
	c := NewCrack(s)
	ic := c.NewICrack()
	return ic.CrackAuthPass()
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
		gologger.Errorf("CrackAuthUser() or CrackAuthPass() is Empty for %s", s)
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

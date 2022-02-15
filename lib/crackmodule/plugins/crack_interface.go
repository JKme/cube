package plugins

type ICrack interface {
	SetName() string   // 插件名称
	Desc() string      //插件作用描述
	Load() bool        // ALL选项的时候，是否加载
	GetPort() string   //设置端口
	Exec() CrackResult //运行task，返回string
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

func (c *Crack) NewCrack() ICrack {
	switch c.Name {
	case "ssh":
		return &SshCrack{c}
	default:
		return &FtpCrack{c}
	}
}

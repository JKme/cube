package plugins

var CRACK_KEYS []string

type ICrack interface {
	SetName() string   // 插件名称
	IsLoad() bool      // ALL选项的时候，是否加载
	SetPort() string   //设置端口
	Exec() CrackResult //运行task，返回string
}

func AddKeys(s string) {
	CRACK_KEYS = append(CRACK_KEYS, s)
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

func (c *Crack) GetPort() string {
	ic := c.NewICrack()
	return ic.SetPort()
}

func (c *Crack) GetLoadStatus() string {
	ic := c.NewICrack()
	if ic.IsLoad() == true {
		return "Y"
	} else {
		return "N"
	}
}

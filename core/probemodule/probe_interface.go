package probemodule

type Probe struct {
	Ip   string
	Port string
	Name string
}

type ProbeResult struct {
	Probe  Probe
	Result string
	Err    error
}

var ProbeKeys []string

type IProbe interface {
	ProbeName() string        //插件名称
	ProbePort() string        //默认端口
	ProbeSkipPortCheck() bool //是否TCP协议
	ProbeExec() ProbeResult   //执行插件
}

func AddProbeKeys(s string) {
	ProbeKeys = append(ProbeKeys, s)
}

func (p *Probe) NewIProbe() IProbe {
	switch p.Name {
	case "oxid":
		return &Oxid{p}
	case "smb":
		return &Smb{p}
	default:
		return nil
	}
}

func NewProbe(s string) Probe {
	return Probe{
		Name: s,
	}
}

func GetName(s string) string {
	c := NewProbe(s)
	ic := c.NewIProbe()
	return ic.ProbeName()
}

func GetPort(s string) string {
	c := NewProbe(s)
	ic := c.NewIProbe()
	return ic.ProbePort()
}

//func GetLoadStatus(s string) bool {
//	c := NewProbe(s)
//	ic := c.NewIProbe()
//	return ic.ProbeLoad()
//}

func GetTCP(s string) bool {
	c := NewProbe(s)
	ic := c.NewIProbe()
	return ic.ProbeSkipPortCheck()
}

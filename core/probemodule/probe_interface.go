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
	ProbeName() string      //插件名称
	ProbePort() string      //默认端口
	PortCheck() bool        //是否需要端口检查
	ProbeExec() ProbeResult //执行插件
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
	case "docker":
		return &Docker{p}
	case "dubbo":
		return &Dubbo{p}
	case "netbios":
		return &Netbios{p}
	case "ms17010":
		return &Ms17010{p}
	case "mssql":
		return &Mssql{p}
	case "ping":
		return &Ping{p}
	case "rmi":
		return &Rmi{p}
	case "smbghost":
		return &SmbGhost{p}
	case "winrm":
		return &Winrm{p}
	case "wmi":
		return &Wmi{p}
	case "zookeeper":
		return &Zookeeper{p}
	case "etcd":
		return &Etcd{p}
	case "k8s":
		return &K8s{p}
	case "jboss":
		return &JBoss{p}
	case "prometheus":
		return &Prometheus{p}
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

func GetProbePort(s string) string {
	c := NewProbe(s)
	ic := c.NewIProbe()
	return ic.ProbePort()
}

//func GetLoadStatus(s string) bool {
//	c := NewProbe(s)
//	ic := c.NewIProbe()
//	return ic.ProbeLoad()
//}

func IsPortCheck(s string) bool {
	c := NewProbe(s)
	ic := c.NewIProbe()
	return ic.PortCheck()
}

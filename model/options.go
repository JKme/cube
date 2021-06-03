package model

type GlobalOptions struct {
	Threads int
	Timeout int
	Verbose bool
	Output  string
	//Plugin  string
}

type ProbeOptions struct {
	Target     string
	TargetFile string
	Port       int
	ScanPlugin string
}

type Service struct {
	Schema string
	Ip     string
	Port   int
}

type SqlcmdOptions struct {
	Service  string
	User     string
	Password string
	Query    string
}

func NewGlobalOptions() *GlobalOptions {
	return &GlobalOptions{}
}

func NewProbeOptions() *ProbeOptions {
	return &ProbeOptions{}
}

func NewSqlcmdOptions() *SqlcmdOptions {
	return &SqlcmdOptions{}
}

type CrackOptions struct {
	Ip       string
	IpFile   string
	User     string
	UserFile string
	Pass     string
	PassFile string
	Port     string
}

func NewCrackOptions() *CrackOptions {
	return &CrackOptions{}
}

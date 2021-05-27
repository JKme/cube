package model

type GlobalOptions struct {
	Threads int
	Timeout int
	Verbose bool
	Output  string
	Plugin  string
	//Plugin  string
}

type ProbeOptions struct {
	Target     string
	TargetFile string
	Port       int
	ScanPlugin string
}

func NewGlobalOptions() *GlobalOptions {
	return &GlobalOptions{}
}

func NewProbeOptions() *ProbeOptions {
	return &ProbeOptions{}
}

type crackOptions struct {
	Target     string
	TargetFile string
	User       string
	UserFile   string
	Pass       string
	PassFile   string
	Port       int
}

type sqlcmdOptions struct {
	Ip    string
	Port  int
	Query string
}

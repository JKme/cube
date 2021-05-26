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
	ScanFunc   string
}

func NewGlobalOptions() *GlobalOptions {
	return &GlobalOptions{}
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

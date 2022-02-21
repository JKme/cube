package model

type GlobalOptions struct {
	Threads int
	Timeout int
	Verbose bool
	Output  string
	Delay   int
}

type ProbeOptions struct {
	Target     string
	TargetFile string
	Port       string
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

func NewProbeOptions() *ProbeOptions {
	return &ProbeOptions{}
}

func NewSqlcmdOptions() *SqlcmdOptions {
	return &SqlcmdOptions{}
}

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

type SqlcmdOptions struct {
	Ip           string
	Port         int
	User         string
	Password     string
	SqlcmdPlugin string
	Query        string
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

type crackOptions struct {
	Target     string
	TargetFile string
	User       string
	UserFile   string
	Pass       string
	PassFile   string
	Port       int
}

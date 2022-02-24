package model

type GlobalOptions struct {
	Threads int
	Timeout int
	Verbose bool
	Output  string
	Delay   int
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

func NewSqlcmdOptions() *SqlcmdOptions {
	return &SqlcmdOptions{}
}

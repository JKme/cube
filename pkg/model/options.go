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

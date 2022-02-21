package core

type GlobalOptions struct {
	Threads int
	Timeout int
	Verbose bool
	Output  string
	Delay   int
}

func NewGlobalOptions() *GlobalOptions {
	return &GlobalOptions{}
}

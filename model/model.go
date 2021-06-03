package model

type ProbeTask struct {
	Ip         string
	Port       int
	ScanPlugin string
}

type ProbeTaskResult struct {
	ProbeTask ProbeTask
	Result    string
	Err       error
}

type CrackTask struct {
	Ip          string
	Port        string
	User        string
	Password    string
	CrackPlugin string
}

type CrackTaskResult struct {
	CrackTask CrackTask
	Result    string
	Err       error
}

type SqlcmdTask struct {
	Ip           string
	Port         int
	User         string
	Password     string
	SqlcmdPlugin string
	Query        string
}

type SqlcmdTaskResult struct {
	SqlcmdTask SqlcmdTask
	Result     string
	Err        error
}

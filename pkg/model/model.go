package model

type ProbeTask struct {
	Ip         string
	Port       string
	ScanPlugin string
}

type ProbeTaskResult struct {
	ProbeTask ProbeTask
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

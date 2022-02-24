package model

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

package model

type globalOptions struct {
	Threads		int
	Timeout		int
	Verbose 	bool
	Output		string
}


type probeOptions struct {
	Target		string
	TargetFile	string
	Port		int
}

type crackOptions struct {
	Target		string
	TargetFile	string
	User		string
	UserFile	string
	Pass		string
	PassFile	string
	Port		int
}

type sqlcmdOptions struct {
	Ip 		string
	Port	int
	Query 	string
}
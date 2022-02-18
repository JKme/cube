package crackmodule

type CrackOptions struct {
	Ip          string
	IpFile      string
	User        string
	UserFile    string
	Pass        string
	PassFile    string
	Port        string
	CrackPlugin string
}

func NewCrackOptions() *CrackOptions {
	return &CrackOptions{}
}

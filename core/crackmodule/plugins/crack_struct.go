package plugins

type Auth struct {
	User     string
	Password string
}
type Crack struct {
	Ip   string
	Port string
	Auth Auth
	Name string
}

type CrackResult struct {
	Crack  Crack
	Result string
	Err    error
}

type CrackOptions struct {
	Ip              string
	IpFile          string
	User            string
	UserFile        string
	Pass            string
	PassFile        string
	Port            string
	CrackPluginName string
}

func NewCrackOptions() *CrackOptions {
	return &CrackOptions{}
}

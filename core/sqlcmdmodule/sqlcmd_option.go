package sqlcmdmodule

type Sqlcmd struct {
	Ip         string
	Port       string
	User       string
	Password   string
	Query      string
	PluginName string
}

type SqlcmdResult struct {
	Sqlcmd Sqlcmd
	Result string
	Err    error
}

type SqlcmdOption struct {
	Ip         string
	Port       string
	User       string
	Password   string
	Query      string
	PluginName string
}

func NewSqlcmdOption() *SqlcmdOption {
	return &SqlcmdOption{}
}

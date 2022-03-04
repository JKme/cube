package sqlcmdmodule

type Mysql struct {
	*Sqlcmd
}

func (m Mysql) SqlcmdName() string {
	return "mysql"
}

func (m Mysql) SqlcmdPort() string {
	return "3306"
}

func (m Mysql) SqlcmdExec() SqlcmdResult {
	//TODO implement me
	panic("implement me")
}

package sqlcmdmodule

type ISqlcmd interface {
	SqlcmdName() string //设定名称
	SqlcmdPort() string
	SqlcmdExec() SqlcmdResult
}

func (s *Sqlcmd) NewISqlcmd() ISqlcmd {
	switch s.Name {
	case "mysql":
		return &Mysql{s}
	case "ssh":
		return &CmdSsh{s}
	default:
		return nil
	}
}

func NewSqlcmd(s string) Sqlcmd {
	return Sqlcmd{
		Name: s,
	}
}

func GetSqlcmdPort(s string) string {
	c := NewSqlcmd(s)
	ic := c.NewISqlcmd()
	return ic.SqlcmdPort()
}

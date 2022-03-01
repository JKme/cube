package sqlcmdmodule

type ISqlcmd interface {
	SqlcmdName() string //设定名称
	SqlcmdPort() string
	SqlcmdExec() SqlcmdResult
}

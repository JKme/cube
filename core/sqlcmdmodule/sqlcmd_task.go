package sqlcmdmodule

import (
	"cube/core"
	"cube/gologger"
)

func StartSqlcmd(opt *SqlcmdOption, globalopt *core.GlobalOption) {
	gologger.Info(opt)
}

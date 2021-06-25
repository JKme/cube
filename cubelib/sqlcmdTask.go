package cubelib

import (
	"cube/log"
	"cube/model"
	Plugins "cube/plugins"
	"fmt"
	"strings"
)

func SaveSqlcmdReport(taskResult model.SqlcmdTaskResult) {
	if len(taskResult.Result) > 0 {
		s := fmt.Sprintf("[->>>>>]: %s\n[->>>>>]: %s:%d\n", taskResult.SqlcmdTask.SqlcmdPlugin, taskResult.SqlcmdTask.Ip, taskResult.SqlcmdTask.Port)
		s1 := fmt.Sprintf("[output]:\n%s", taskResult.Result)
		fmt.Println(s + s1)
	}
}

func StartSqlcmdTask(opt *model.SqlcmdOptions, globalopts *model.GlobalOptions) {
	s, err := ParseService(opt.Service)
	if err != nil {
		log.Error(err)
	}

	_, key := Plugins.SqlcmdFuncMap[s.Schema]
	if !key {
		log.Errorf("%s plugin not found, available plugins: %s", s.Schema, strings.Join(Plugins.SqlcmdKeys, ","))
	}

	task := model.SqlcmdTask{Ip: s.Ip, Port: s.Port, User: opt.User, Password: opt.Password, SqlcmdPlugin: s.Schema, Query: opt.Query}
	fn := Plugins.SqlcmdFuncMap[task.SqlcmdPlugin]
	SaveSqlcmdReport(fn(task))
}

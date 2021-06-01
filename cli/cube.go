package cli

import (
	"cube/cubelib"
	"cube/log"
	"cube/model"
	Plugins "cube/plugins"
	"cube/util"
	"strings"
)

func StartProbeTask(opt *model.ProbeOptions, globalopts *model.GlobalOptions) {
	ips, err := util.ParseIP(opt.Target, opt.TargetFile)
	if err != nil {
		log.Error(err)
	}
	pluginList, err := cubelib.ValidPlugin(opt.ScanPlugin)
	if err != nil {
		log.Fatal(err)
	}
	if !cubelib.Subset(pluginList, Plugins.ProbeKeys) {
		log.Fatalf("plugins not found: %s", pluginList)
	}

	tasks := cubelib.GenerateTasks(ips, opt.Port, pluginList)
	cubelib.RunTasks(tasks, globalopts.Threads, globalopts.Timeout)
}

func StartSqlcmdTask(opt *model.SqlcmdOptions, globalopts *model.GlobalOptions) {
	s, err := cubelib.ParseService(opt.Service)
	if err != nil {
		log.Fatal(err)
	}

	_, key := Plugins.SqlcmdFuncMap[s.Schema]
	if !key {
		log.Fatalf("Available Plugins: %s", strings.Join(Plugins.SqlcmdKeys, ","))
	}

	task := model.SqlcmdTask{Ip: s.Ip, Port: s.Port, User: opt.User, Password: opt.Password, SqlcmdPlugin: s.Schema, Query: opt.Query}
	fn := Plugins.SqlcmdFuncMap[task.SqlcmdPlugin]
	cubelib.SaveSqlcmdReport(fn(task))
}

package interfaces

import "cube/pkg/model"

type CrackPluginInterface interface {
	Name() string // 插件名称
	Desc() string //插件作用描述
	Load() bool // ALL选项的时候，是否加载
	SetPort() string  //设置端口
	Exec(model.CrackTask) string  //运行task，返回string
}


type Result interface {
	ResultToString() (string, error)  //probe、crack、sqlcmd都实现获取结果的接口
}
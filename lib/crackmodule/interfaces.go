package crackmodule

import (
	"cube/lib/crackmodule/plugins"
)

type CrackPluginInterface interface {
	Name() string // 插件名称
	Desc() string //插件作用描述
	Load() bool // ALL选项的时候，是否加载
	Port() string  //设置端口
	Exec(plugins.Crack) plugins.CrackResult  //运行task，返回string
}


type Result interface {
	ResultToString() (string, error)  //probe、crack、sqlcmd都实现获取结果的接口
}

var CrackPlugins map[string]CrackPluginInterface

func init(){
	CrackPlugins = make(map[string]CrackPluginInterface)
}

func RegisterCrackPlugin(name string, crackPluginInterface CrackPluginInterface){
	CrackPlugins[name] = crackPluginInterface
}

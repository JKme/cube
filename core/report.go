package core

type IReport interface {
	Module() string //三个模块，每个模块一个名称
	Save()          //每个模块实现保存结果到csv的方法
}

//type Result interface {
//	ResultToString() (string, error)  //probe、crack、sqlcmd都实现获取结果的接口
//}
//

//var ICrackMap map[string]ICrack
//
//func init() {
//	ICrackMap = make(map[string]ICrack)
//}

package cubelib

import (
	"cube/model"
	Plugins "cube/plugins"
	"fmt"
	"strings"
	"sync"
)

func executeSqlcmdTask(task model.SqlcmdTask, wg *sync.WaitGroup) {
	defer wg.Done()

	//fmt.Println("Hello")
	fn := Plugins.SqlcmdFuncMap[task.CrackTask.CrackPlugin]
	saveSqlcmdReport(fn(task))

}

func saveSqlcmdReport(taskResult model.SqlcmdTaskResult) {
	if len(taskResult.Result) > 0 {
		fmt.Println(strings.Repeat("=", 20))
		fmt.Printf("%s:\n%s", taskResult.SqlcmdTask.CrackTask.Ip, taskResult.Result)
		fmt.Println(strings.Repeat("=", 20))
	}
}

package cubelib

import (
	"cube/model"
	Plugins "cube/plugins"
	"fmt"
	"sync"
)

func executeSqlcmdTask(task model.SqlcmdTask, wg *sync.WaitGroup) {
	defer wg.Done()

	//fmt.Println("Hello")
	fn := Plugins.SqlcmdFuncMap[task.SqlcmdPlugin]
	SaveSqlcmdReport(fn(task))

}

func SaveSqlcmdReport(taskResult model.SqlcmdTaskResult) {
	if len(taskResult.Result) > 0 {
		s := fmt.Sprintf("[*]: %s\n[*]: %s:%d\n", taskResult.SqlcmdTask.SqlcmdPlugin, taskResult.SqlcmdTask.Ip, taskResult.SqlcmdTask.Port)
		s1 := fmt.Sprintf("[output]:\n %s", taskResult.Result)
		fmt.Println(s + s1)
	}
}

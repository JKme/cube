package util

import (
	"cube/model"
	"fmt"
	"strings"
)

func saveReport(taskResult model.TaskResult){
	if len(taskResult.Result) > 0 {
		fmt.Println(strings.Repeat("=", 20))
		fmt.Printf("%s:\n%s", taskResult.Task.Ip, taskResult.Result)
		fmt.Println(strings.Repeat("=", 20))
	}
}
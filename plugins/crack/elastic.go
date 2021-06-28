package crack

import (
	"cube/model"
	"cube/util"
	"fmt"
	"gopkg.in/olivere/elastic.v3"
)

func ElasticCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	client, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%v:%v", task.Ip, task.Port)),
		elastic.SetBasicAuth(task.Auth.User, task.Auth.Password),
	)
	if err == nil {
		_, _, err = client.Ping(fmt.Sprintf("http://%v:%v", task.Ip, task.Port)).Do()
		if err == nil {
			result.Result = util.Green(fmt.Sprintf("User: %s\tPassword: %s \t", task.Auth.User, task.Auth.Password))
		}
	}
	return result
}

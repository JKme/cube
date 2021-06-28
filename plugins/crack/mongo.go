package crack

import (
	"cube/model"
	"cube/util"
	"fmt"
	"gopkg.in/mgo.v2"
)

func MongoCrack(task model.CrackTask) (result model.CrackTaskResult) {
	result = model.CrackTaskResult{CrackTask: task, Result: "", Err: nil}
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", task.Auth.User, task.Auth.Password, task.Ip, task.Port, "test")
	session, err := mgo.DialWithTimeout(url, model.ConnectTimeout)

	if err == nil {
		defer session.Close()
		err = session.Ping()
		if err == nil {
			result.Result = util.Green(fmt.Sprintf("User: %s\tPassword: %s \t", task.Auth.User, task.Auth.Password))
		}
	}
	return result
}

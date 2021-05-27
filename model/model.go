package model

type ProbeTask struct {
	Ip         string
	Port       int
	ScanPlugin string
}

type ProbeTaskResult struct {
	ProbeTask ProbeTask
	Result    string
	Err       error
}

//type ScanFunc func(task Task) (err error, taskResult TaskResult)

var (
	ScanPluginMapPort map[string]int
)

func init() {

	ScanPluginMapPort = make(map[string]int)
	ScanPluginMapPort["OXID"] = 135
	ScanPluginMapPort["SMB"] = 445

}

package core

import "cube/gologger"

type GlobalOption struct {
	Threads int
	Timeout int
	Verbose bool
	Output  string
	Delay   float64
}

func NewGlobalOptions() *GlobalOption {
	return &GlobalOption{}
}

func RandomDelay(float float64) float64 {
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	//r1 := r.Float64() * float
	//r1, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", r1), 64) //保留两位小数点
	//if r1 > 0 {
	//	gologger.Debugf("thread is going to sleep %v second", r1)
	//}
	//return r1
	gologger.Debugf("thread is going to sleep %v second", float)
	return float
}

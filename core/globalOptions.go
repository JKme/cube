package core

import (
	"cube/gologger"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type GlobalOptions struct {
	Threads int
	Timeout int
	Verbose bool
	Output  string
	Delay   float64
}

func NewGlobalOptions() *GlobalOptions {
	return &GlobalOptions{}
}

func SleepDelay(float float64) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r1 := r.Float64() * float
	r1, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", r1), 64) //保留两位小数点
	gologger.Debugf("Sleep %v Second", r1)
	time.Sleep(time.Duration(r1) * time.Millisecond * 1000)
}

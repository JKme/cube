package report

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrentSlice_Append(t *testing.T) {
	c := CsvCell{
		Ip:     "127",
		Module: "crack",
		Cell:   "123456",
	}
	c1 := CsvCell{
		Ip:     "888",
		Module: "Probe",
		Cell:   "999",
	}
	cs := ConcurrentSlice{
		RWMutex: sync.RWMutex{},
		items:   nil,
	}
	cs.Append(c)
	cs.Append(c)
	cs.Append(c1)

	for k := range cs.Iter() {
		fmt.Println(k.Index, k.Value.Ip)
	}
}

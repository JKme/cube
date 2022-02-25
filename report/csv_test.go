package report

import (
	"fmt"
	"sort"
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

type Member struct {
	Id        int
	LastName  string
	FirstName string
}

func SortByLastNameAndFirstName(members []Member) {
	sort.SliceStable(members, func(i, j int) bool {
		mi, mj := members[i], members[j]
		switch {
		case mi.LastName != mj.LastName:
			return mi.LastName < mj.LastName
		default:
			return mi.FirstName < mj.FirstName
		}
	})
}

func TestSort(t *testing.T) {
	members := []Member{
		{0, "The", "quick"},
		{1, "brown", "fox"},
		{2, "jumps", "over"},
		{3, "brown", "grass"},
		{4, "brown", "grass"},
		{5, "brown", "grass"},
		{6, "brown", "grass"},
		{7, "brown", "grass"},
		{8, "brown", "grass"},
		{9, "brown", "grass"},
		{10, "brown", "grass"},
		{11, "brown", "grass"},
	}

	SortByLastNameAndFirstName(members)

	for _, member := range members {
		fmt.Println(member)
	}
}

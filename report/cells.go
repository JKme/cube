package report

//http://dnaeon.github.io/concurrent-maps-and-slices-in-go/
import "sync"

type CsvCell struct {
	Ip     string
	Module string
	Cell   string
}

type ConcurrentSlice struct {
	sync.RWMutex
	items []CsvCell
}

type ConcurrentSliceItem struct {
	Index int
	Value CsvCell
}

var ConcurrentSlices ConcurrentSlice
var CsvShells []CsvCell

func (cs *ConcurrentSlice) Append(item CsvCell) {
	cs.Lock()
	defer cs.Unlock()

	cs.items = append(cs.items, item)
}

func (cs *ConcurrentSlice) Iter() <-chan ConcurrentSliceItem {
	c := make(chan ConcurrentSliceItem)

	f := func() {
		cs.Lock()
		defer cs.Unlock()
		for index, value := range cs.items {
			c <- ConcurrentSliceItem{index, value}
		}
		close(c)
	}
	go f()

	return c
}

//https://juejin.cn/post/6844904134592692231
//https://stackoverflow.com/questions/36122668/how-to-sort-struct-with-multiple-sort-parameters

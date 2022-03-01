package report

import "sort"

type KV struct {
	Key   string
	Value int
}

func SortSlice(ss []KV) {
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
}

func SortPlug(css []CsvCell) []KV {
	var srMap = make(map[string]int)
	for _, v := range css {
		sr := v.Module
		srMap[sr] += 1
	}
	var ss []KV
	for k, v := range srMap {
		ss = append(ss, KV{
			Key:   k,
			Value: v,
		})
	}
	SortSlice(ss)
	return ss
}

func SortIP(css []CsvCell) []KV {
	var srMap = make(map[string]int)
	for _, v := range css {
		sr := v.Ip
		srMap[sr] += 1
	}
	var ss []KV
	for k, v := range srMap {
		ss = append(ss, KV{
			Key:   k,
			Value: v,
		})
	}
	SortSlice(ss)
	return ss
}

func GetKeys(kvs []KV) (l []string) {
	for _, data := range kvs {
		l = append(l, data.Key)
	}
	return l
}

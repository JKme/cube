package report

import (
	"cube/gologger"
	"fmt"
	"github.com/xuri/excelize/v2"
	"sort"
	"strconv"
	"strings"
)

type KV struct {
	Key   string
	Value int
}

func SortSlice(ss []KV) {
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
}

func SortPlug() []KV {
	var srMap = make(map[string]int)
	for k := range ConcurrentSlices.Iter() {
		sr := k.Value.Module
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

func SortIP() []KV {
	var srMap = make(map[string]int)
	for k := range ConcurrentSlices.Iter() {
		sr := k.Value.Ip
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

func convertXY(x, y int) string {
	return fmt.Sprintf(string(rune(x)) + strconv.Itoa(y))
}

func inc(i int) int {
	i += 1
	return i
}

func GetCsvShellValue(ip, module string, csvs []CsvCell) (s string) {
	//传入module 和 ip获取到string
	for _, csv := range csvs {
		if csv.Ip == ip && csv.Module == module {
			s = csv.Cell
			break
		} else {
			s = ""
		}
	}
	return s
}

func ExportTitle(heads []string, ips []string, csvs []CsvCell) {
	excel := excelize.NewFile()
	_ = excel.SetSheetRow("Sheet1", "A1", &heads)

	style, err := excel.NewStyle(`{
					"border":[
                                {
                                        "type":"left",
                                        "color":"000000",
                                        "style":1
                                },
                                {
                                        "type":"top",
                                        "color":"000000",
                                        "style":1
                                },
                                {
                                        "type":"bottom",
                                        "color":"000000",
                                        "style":1
                                },
                                {
                                        "type":"right",
                                        "color":"000000",
                                        "style":1
                                }
                        ],
            "alignment":{
                "wrap_text":true
            }
        }`)
	//"horizontal":"left",
	//	"vertical":"left",
	if err != nil {
		gologger.Error(err)
	}

	y := 2
	for _, ip := range ips {
		x := 65
		excel.SetCellStyle("Sheet1", fmt.Sprintf("A%d", y), fmt.Sprintf("A%d", y), style)
		excel.SetCellValue("Sheet1", fmt.Sprintf("A%d", y), ip)
		x += 1
		for _, plug := range heads[1:] {
			data := GetCsvShellValue(ip, plug, csvs)
			if len(data) > 0 {
				excel.SetCellStyle("Sheet1", convertXY(x, y), convertXY(x, y), style)
				excel.SetCellValue("Sheet1", convertXY(x, y), strings.Trim(data, " "))
				x += 1
			} else {
				excel.SetCellStyle("Sheet1", convertXY(x, y), convertXY(x, y), style)
				x += 1
			}
		}
		y += 1
	}

	if err := excel.SaveAs("Book1.xlsx"); err != nil {
		gologger.Error(err)
	}

	//axis := fmt.Sprintf("A%d", key)
	//tmp,_ := datum.([]interface{})
	//if k.Index
	//_ = excel.SetSheetRow("Sheet1", axis)
}

//k += 1

func ExportIP(ips []string) {

}

//

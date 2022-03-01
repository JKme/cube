package report

import (
	"fmt"
	"github.com/xuri/excelize/v2"
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

func WriteXlsx() {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	//if err := f.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`); err != nil {
	//	println(err.Error())
	//	return
	//}
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		println(err.Error())
	}
}

func TestWriteXlsx(t *testing.T) {
	WriteXlsx()
}

func TestSortMap(t *testing.T) {
	var kvs []KV
	kvs = append(kvs, KV{
		Key:   "Probe_smb",
		Value: 14,
	})
	kvs = append(kvs, KV{
		Key:   "Probe_oxid",
		Value: 13,
	})
	fmt.Println(GetKeys(kvs))
}

//func TestExport(t *testing.T) {
//	l := []string{"Probe", "Hello"}
//}

func Xls2() {
	l := []string{"IP", "127", "198", "133"}
	f := excelize.NewFile()
	// 创建一个工作表
	_ = f.SetSheetRow("Sheet1", "A1", &l)

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func TestXls2(t *testing.T) {
	//Xls2()
	//c := 'A'
	//asciiValue := int(c)
	//fmt.Println(asciiValue)
	//B := asciiValue + 1
	//fmt.Println(string(rune(B)))

	fmt.Println(convertXY(66, 2))
}

func TestExportTitle(t *testing.T) {
	heads := []string{"Probe_smb", "Probe_oxid"}
	//ips := []string{"127.0.0.1", "10.10.10.1", "10.10.10.2"}
	var cvs []CsvCell
	heads = append([]string{"IP"}, heads...)

	CsvCell0 := CsvCell{
		Ip:     "127.0.0.1",
		Module: "Probe_smb",
		Cell:   "smb222",
	}
	CsvCell1 := CsvCell{
		Ip:     "10.10.10.1",
		Module: "Probe_smb",
		Cell:   "smb",
	}
	CsvCell2 := CsvCell{
		Ip:     "127.0.0.1",
		Module: "Probe_oxid",
		Cell:   "oxid3333",
	}
	CsvCell3 := CsvCell{
		Ip:     "10.10.10.2",
		Module: "Probe_oxid",
		Cell:   "oxid3333",
	}

	CsvCell4 := CsvCell{
		Ip:     "10.10.10.2",
		Module: "Probe_oxid",
		Cell:   "oxid3333",
	}

	cvs = append(cvs, CsvCell0)
	cvs = append(cvs, CsvCell1)
	cvs = append(cvs, CsvCell2)
	cvs = append(cvs, CsvCell3)
	cvs = append(cvs, CsvCell4)
	//ExportTitle(heads, ips, cvs)

	r := RemoveDuplicateCSS(cvs)
	fmt.Println(r)
}

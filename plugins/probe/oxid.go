package probe

import (
	"bytes"
	"cube/model"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net"
)

func OxidProbe(task model.ProbeTask) (result model.ProbeTaskResult) {
	result = model.ProbeTaskResult{ProbeTask: task, Result: "", Err: nil}
	//log.Printf("Oxid Running Debug: %s", task.Ip)
	dl := net.Dialer{Timeout: model.ConnectTimeout}
	t := fmt.Sprintf("%s:%d", task.Ip, model.CommonPortMap[task.ScanPlugin])
	conn, err := dl.Dial("tcp", t)

	// defer conn.Close()
	if err != nil {
		result.Err = err
		//log.Printf("Oxid Running Error: %s:%s", task.Ip, err)
		return result
	}

	conn.Write([]byte("\x05\x00\x0b\x03\x10\x00\x00\x00\x48\x00\x00\x00\x01\x00\x00\x00\xb8\x10\xb8\x10\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x01\x00\xc4\xfe\xfc\x99\x60\x52\x1b\x10\xbb\xcb\x00\xaa\x00\x21\x34\x7a\x00\x00\x00\x00\x04\x5d\x88\x8a\xeb\x1c\xc9\x11\x9f\xe8\x08\x00\x2b\x10\x48\x60\x02\x00\x00\x00"))

	tmpByte := make([]byte, 1024)
	conn.Read(tmpByte)

	// dcerpc finish

	// IOXIDResolve start
	conn.Write([]byte("\x05\x00\x00\x03\x10\x00\x00\x00\x18\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00"))

	r := make([]byte, 4096)
	_, err = conn.Read(r)
	//if err != nil {
	//	return err
	//}
	/*
		Two parts:
			1. Distributed Computing Enviroment / Remote Prodedure Call Response
			2. DCOM OXID Resolver <- what we need
	*/

	r = r[24+12+2+2:]
	index := bytes.Index(r, []byte("\x09\x00\xff\xff\x00\x00"))
	//if index == -1 {
	//	return errors.New("Not Found")
	//}
	r = r[:index]
	//fmt.Println("hello")
	//fmt.Printf("%v", r)
	//s := string(r)
	////for _, r := range s {
	////	fmt.Println(r)
	////}
	//rs := []rune(s)
	//fmt.Println(rs)
	//r1, _ := gbkToUtf8(r)
	//fmt.Println(r1)
	//fmt.Println(string(r1))
	//https://play.golang.org/p/Q93aIamc_rG
	//https://play.golang.org/p/M7sc7cGv4l3
	//http://cs50mu.github.io/blog/2019/05/19/a-encoding-problem-in-golang/
	//https://github.com/animesh-server-dot-files/go/blob/381c0dd07cb51c5607b4a8e66b814292f2225fd6/v1.16.4/source/src/unicode/utf8/example_test.go
	//results := []string{}
	var results []string

	for {
		if len(r) == 0 {
			break
		}
		index = bytes.Index(r, []byte("\x00\x00\x00"))
		results = append(results, dataGet(r[:index+3]))
		r = r[index+3:]
	}
	if len(results) > 0 {
		var s string
		for _, v := range results {
			s = s + fmt.Sprintf("\t"+v+"\n")
		}

		result.Result = s
	}
	archPayload := []byte{ /* Packet 186 */
		0x05, 0x00, 0x0b, 0x03, 0x10, 0x00, 0x00, 0x00,
		0x48, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0xb8, 0x10, 0xb8, 0x10, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00,
		0x08, 0x83, 0xaf, 0xe1, 0x1f, 0x5d, 0xc9, 0x11,
		0x91, 0xa4, 0x08, 0x00, 0x2b, 0x14, 0xa0, 0xfa,
		0x03, 0x00, 0x00, 0x00, 0x33, 0x05, 0x71, 0x71,
		0xba, 0xbe, 0x37, 0x49, 0x83, 0x19, 0xb5, 0xdb,
		0xef, 0x9c, 0xcc, 0x36, 0x01, 0x00, 0x00, 0x00}
	return result
}

func dataGet(data []byte) string {
	if bytes.HasPrefix(data, []byte("\x07\x00")) {
		return string(data[:len(data)-3])
	}
	return ""
}

func gbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return d, nil
}

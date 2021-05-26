package probe

import (
	"bytes"
	probe2 "cube/cubelib/probe"
	"cube/model"
	"strings"
	//"errors"
	"fmt"
	"net"
)

func OxidProbe(task probe2.Task) (result probe2.TaskResult) {
	result = probe2.TaskResult{Task: task, Result: "", Err: nil}
	//log.Printf("Oxid Running Debug: %s", task.Ip)
	dl := net.Dialer{Timeout: model.TIMEOUT}
	t := fmt.Sprintf("%s:%d", task.Ip, model.ScanPluginMapPort[task.ScanPlugin])
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
		//fmt.Println(results)
		//r1 := fmt.Sprintf("[+] %s\n", task.Ip)
		//for _, v := range results {
		//	fmt.Println("\t" + v)
		//
		//}
		//for _, v := range results {
		//	fmt.Println("\t" + v)
		//}
		//fmt.Println(results)
		result.Result = strings.Join(results, "\n")
	}
	return result
}

func dataGet(data []byte) string {
	if bytes.HasPrefix(data, []byte("\x07\x00")) {
		return string(data[:len(data)-3])
	}
	return ""
}

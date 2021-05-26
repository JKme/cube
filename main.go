package main

import "cube/cmd"

func main() {
	//var (
	//	scanPlugin  	string
	//	scanTargets		string
	//	scanTargetsFile	string
	//	scanPort		int
	//	timeout     	int
	//	scanNum			int
	//	verbose         bool
	//)
	//flag.StringVar(&scanPlugin, "m", "smb", "scan plugins")
	//flag.StringVar(&scanTargets, "ip", "", "scan targets")
	//flag.StringVar(&scanTargetsFile, "hf", "", "host file, -hs ip.txt")
	//flag.IntVar(&scanPort, "p", 445,"scan port")
	//flag.IntVar(&timeout, "t", 5, "scan timeout")
	//flag.IntVar(&scanNum, "n", 10, "scan threads")
	//flag.BoolVar(&verbose, "vv", false, "verbose")
	//
	//flag.Parse()
	////fmt.Println(scanPlugin, scanTargets, scanPort, timeout, scanNum)
	//if verbose {
	//	log.InitLog("debug")
	//} else {
	//	log.InitLog("error")
	//}
	//util.Scan(scanPlugin, scanTargets,scanTargetsFile, scanPort, timeout, scanNum)
	cmd.Execute()
}

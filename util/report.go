package util

import (
	"cube/model"
	"fmt"
	"strings"
	"unicode"
)

func saveReport(taskResult model.ProbeTaskResult) {
	if len(taskResult.Result) > 0 {
		fmt.Println(strings.Repeat("=", 20))
		fmt.Printf("%s:\n%s", taskResult.ProbeTask.Ip, taskResult.Result)
		fmt.Println(strings.Repeat("=", 20))
	}
}

var (
	Info   = Teal_
	Yellow = Yellow_
	Red    = Red_
	Green  = Green_
)

var (
	Black_   = Color("\033[1;30m%s\033[0m")
	Red_     = Color("\033[1;31m%s\033[0m")
	Green_   = Color("\033[1;32m%s\033[0m")
	Yellow_  = Color("\033[1;33m%s\033[0m")
	Purple_  = Color("\033[1;34m%s\033[0m")
	Magenta_ = Color("\033[1;35m%s\033[0m")
	Teal_    = Color("\033[1;36m%s\033[0m")
	White_   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func StrXor(message string, keywords string) string {
	messageLen := len(message)
	keywordsLen := len(keywords)

	result := ""

	for i := 0; i < messageLen; i++ {
		result += string(message[i] ^ keywords[i%keywordsLen])
	}
	return result
}

package util

func Contains(str string, slice []string) bool {
	//str是否在slice列表里面
	for _, value := range slice {
		if str == value {
			return true
		}
	}
	return false
}

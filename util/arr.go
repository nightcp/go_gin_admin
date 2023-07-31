package util

type arrUtil struct {
}

var ArrUtil = &arrUtil{}

// InArray 搜索字符串是否在切片中
func (au *arrUtil) InArray(needle string, haystack []string) bool {
	found := false
	for _, v := range haystack {
		if needle == v {
			found = true
			break
		}
	}
	return found
}

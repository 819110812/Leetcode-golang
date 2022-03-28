package _4LongestCommonPrefix

// LongestCommonPrefix TODO: 重做
func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		pre = getPrefix(pre, strs[i])
		if len(pre) == 0 {
			break
		}
	}

	return pre
}

func getPrefix(str1 string, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

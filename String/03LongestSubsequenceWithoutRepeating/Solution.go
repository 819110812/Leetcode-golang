package _3LongestSubsequenceWithoutRepeating

func LongestSubsequenceWithoutRepeating(str string) int {
	length := len(str)
	if length < 2 {
		return length
	}

	var record [127]int
	res, left, right := 0, 0, -1

	//
	for left < length {
		if right+1 < length && record[str[right+1]] == 0 {
			record[str[right+1]]++
			right++
		} else {
			record[str[left]]--
			left++
		}
		res = max(res, right-left+1)
	}

	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

package _3LongestSubsequenceWithoutRepeating

func LongestSubsequenceWithoutRepeating(str string) int {
	//var maxLen int
	var length = len(str)
	if length <= 1 {
		return length
	}

	// 结果长度，左侧边界，右侧边界
	var maxlen, left, right = 0, 0, -1

	var freq [127]int
	//var freq = make(map[byte]int, length)
	// 左侧边界小于字符串长度就说明仍然有字符串需要检查
	// [a] b c
	for left < length {
		// 每次迭代右边界都要往右一位
		// right + 1 表示下一个字符串，如果小于字符串总长度则没有结束
		// freq[str[right+1]] 表示当前字符串中的字符出现的次数, 如果为0则表示没有出现过
		// 当满足字符串没有遍历完，且下一个字符串没有出现过时，更新数据
		if right+1 < length && freq[str[right+1]] == 0 {
			freq[str[right+1]]++
			right++
		} else {
			// 当上一个条件不满足时，说明，要么右侧已经到头，要么已经开始重复了
			// 此时需要开始挪动左侧的边界，并且将挪出的字符减少记录数
			freq[str[left]]--
			left++
		}
		// 比较长度，是否需要更新做大长度
		// right - left + 1 表示当前的长度
		maxlen = max(maxlen, right-left+1)
	}

	return maxlen
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

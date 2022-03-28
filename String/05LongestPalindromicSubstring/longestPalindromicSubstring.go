package _5LongestPalindromicSubstring

func LongestPalindromicSubstring(str string) string {
	length := len(str)
	if length < 2 {
		return str
	}

	// left，right 不断更新的指针
	// pl,pr 用来记录left,right最佳结果的指针
	left, right, pl, pr := 0, -1, 0, 0

	for left < length {
		for right+1 < length && str[right+1] == str[left] {
			right++
		}

		for left-1 >= 0 && right+1 < length && str[left-1] == str[right+1] {
			left--
			right++
		}

		if (right - left) > (pr - pl) {
			pl, pr = left, right
		}

		// 为什么要这么更新
		left = (right+left)/2 + 1
		right = left
	}

	return str[pl : pr+1]
}

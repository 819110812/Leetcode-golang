package _4LongestPalindromicSubstring

func LongestPalindromicSubstring(str string) string {
	if len(str) == 0 {
		return ""
	}

	left, right, pl, pr := 0, -1, 0, 0
	for left < len(str) {
		// 右侧指针挪一位，如果小于总长度，说明没结束
		// 如果是和左侧的字符相等，说明是回文，右侧单右侧挪动
		// 这边是持续挪动，直到不相等
		// test case； bbbb， 这种情况如果不挪动会出现 pl=pr=0
		for right+1 < len(str) && str[left] == str[right+1] {
			right++
		}
		// 如果上面的条件不满足，说明此时右侧的字符不是回文，左侧挪动，并且同时右侧挪动，找到回文的边界在哪里
		// left -1 >=0 :定义左侧边界，使的left--不越界
		// right + 1 < len(str) :定义右侧边界，使得right++不越界
		// str[left-1] == str[right+1] : 如果左侧的字符和右侧的字符相等，说明是回文，左右指针挪动
		for left-1 >= 0 && right+1 < len(str) && str[left-1] == str[right+1] {
			left--
			right++
		}

		// 如果当前的回文长度大于之前的回文长度，则更新
		if right-left > pr-pl {
			//pl = left
			//pr = right
			pl, pr = left, right
		}
		left = (left+right)/2 + 1
		right = left
	}
	return str[pl : pr+1]
}

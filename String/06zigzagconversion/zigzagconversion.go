package _6zigzagconversion

// TODO: 需要复习
//func convert(s string, numRows int) string {
//	if numRows == 1 {
//		return s
//	}
//
//	// matrix 用来放置字符的二维数组
//	// down 初始值为0，意思这个指针在最开始的第一层,用来控制向下的方向
//	// up 初始值为 numRows-2，位于倒数第二行，numRows-1是最后一行的数组表达位置，则numRows-2为数组中的倒数第二行
//	maxtrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
//	for i := 0; i != len(s); {
//		// down 二维数组顺序往下填
//		// 未到下边界
//		if down != numRows {
//			maxtrix[down] = append(maxtrix[down], s[i])
//			down++
//			i++
//		} else if up != 0 {
//			// 已经到了下边界，开始向上挪
//			// up不能超过上边界
//			maxtrix[up] = append(maxtrix[up], s[i])
//			up--
//			i++
//		} else {
//			// 重置指针
//			up = down - 2
//			down = 0
//		}
//
//	}
//	// 收集最终的结果
//	res := make([]byte, 0, len(s))
//	for _, v := range maxtrix {
//		for _, vv := range v {
//			res = append(res, vv)
//		}
//	}
//
//	return string(res)
//}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	down, up := 0, numRows-2
	matrix := make([][]byte, numRows, numRows)

	for i := 0; i < len(s); {
		if down < numRows {
			matrix[down] = append(matrix[down], s[i])
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], s[i])
			up--
			i++
		} else {
			down = 0
			up = numRows - 2
		}
	}

	res := make([]byte, 0)
	for _, row := range matrix {
		for _, col := range row {
			res = append(res, col)
		}
	}
	return string(res)

}

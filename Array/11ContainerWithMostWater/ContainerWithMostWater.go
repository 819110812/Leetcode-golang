package _1ContainerWithMostWater

//11. 盛最多水的容器
//给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//返回容器可以储存的最大水量。
//
//说明：你不能倾斜容器。
// 输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
//示例 2：
//
//输入：height = [1,1]
//输出：1
//

//func MaxArea(height []int) int {
//	if len(height) == 0 {
//		return 0
//	}
//	left, right := 0, len(height)-1
//	res := 0
//	for left < right {
//		area := min(height[left], height[right]) * (right - left)
//		res = max(res, area)
//		if height[left] < height[right] {
//			left++
//		} else {
//			right--
//		}
//	}
//	return res
//}
//
func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		cur := min(height[left], height[right]) * (right - left)
		area = max(area, cur)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}

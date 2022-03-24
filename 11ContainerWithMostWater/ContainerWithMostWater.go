package _1ContainerWithMostWater

func MaxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		res = max(res, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

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

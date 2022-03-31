package _3SearchinRotatedSortedArray

func search(nums []int, target int) int {
	flag, index := contains(nums, target)
	if !flag {
		return -1
	}
	return index
}

func contains(nums []int, target int) (bool, int) {
	if len(nums) == 0 {
		return false, 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return true, i
		}
	}
	return false, 0
}

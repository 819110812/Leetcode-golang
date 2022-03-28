package _210CountHillsandValleysinAnArray

func countHillValley(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	//length := len(nums)
	count := 0
	//flag := 1
	var flag [127]bool

	index := 1
	for index < len(nums)-1 {
		left := index - 1
		right := index + 1
		for left >= 0 && nums[left] == nums[index] {
			left--
		}
		for right < len(nums) && nums[index] == nums[right] {
			right++
		}

		if nums[left] > nums[index] && nums[right] > nums[index] && flag[index-1] == false {
			count++
			flag[index] = true
		}

		if nums[right] < nums[index] && nums[left] < nums[index] && flag[index-1] == false {
			count++
			flag[index] = true
		}
		index++
	}

	return count
}

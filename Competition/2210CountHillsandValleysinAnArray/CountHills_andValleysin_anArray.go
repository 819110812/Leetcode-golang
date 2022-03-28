package _210CountHillsandValleysinAnArray

func countHillValley(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	//length := len(nums)
	count := 0
	//flag := 1

	index := 1
	for index < len(nums)-1 {
		if nums[index] == nums[index-1] {
			index++
			continue
		}

		right := index + 1
		for right < len(nums) && nums[index] == nums[right] {
			right++
		}

		if nums[index] > nums[index-1] {
			if right < len(nums) && nums[index] > nums[right] {
				count++
			}
		}

		if nums[index] < nums[index-1] {
			if right < len(nums) && nums[index] < nums[right] {
				count++
			}
		}

		index++
	}

	return count
}

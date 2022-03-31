package _1NextPermutations

func nextPermutation(nums []int) {
	i, j := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	for i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[i] < nums[j] {
				break
			}
		}
		swap(i, j, &nums)
	}
	reverse(&nums, i+1, len(nums)-1)
}

func swap(i int, j int, nums *[]int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func reverse(nums *[]int, start int, end int) {
	for start < end {
		swap(start, end, nums)
		start++
		end--
	}
}

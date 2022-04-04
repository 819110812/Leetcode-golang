package _4FindFirstandLastPositionofElementinSortedArray

func searchRange(nums []int, target int) []int {
	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false) - 1
	if left > right {
		return []int{-1, -1}
	}
	return []int{left, right}
}

func binarySearch(nums []int, target int, isLeft bool) int {
	left := 0
	right := len(nums) - 1
	ans := len(nums)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target || (isLeft && nums[mid] == target) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}

package _5ThreeSum

import (
	"reflect"
	"sort"
)

// ThreeSum : TODO 重做
func ThreeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < length; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := length - 1
		target := -1 * nums[first]
		for second := first + 1; second < length; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}

		}
	}
	return ans
}

func Contains(parent [][]int, child []int) bool {
	if len(parent[0]) != len(child) {
		return false
	}
	if reflect.DeepEqual(parent, child) {
		return true
	}
	return false
}

func Sum(arr []int) int {
	sums := 0
	for _, num := range arr {
		sums = sums + num
	}
	return sums
}

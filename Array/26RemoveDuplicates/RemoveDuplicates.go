package _6RemoveDuplicates

//func removeDuplicates(nums []int) int {
//	if len(nums) < 2 {
//		return len(nums)
//	}
//	record := map[int]int{}
//
//	index := 0
//	for index < len(nums) {
//		if ok := record[nums[index]]; ok != 0 {
//			nums = append(nums[:index], nums[index+1:]...)
//			index--
//		} else {
//			record[nums[index]] = 1
//		}
//		index++
//	}
//	fmt.Println(record)
//	return len(record)
//}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

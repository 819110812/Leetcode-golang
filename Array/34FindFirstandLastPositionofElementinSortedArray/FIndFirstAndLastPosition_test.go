package _4FindFirstandLastPositionofElementinSortedArray

import "testing"

func Test_Should_Return_Correct_Answer(t *testing.T) {
	res := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	if res[0] != 3 || res[1] != 4 {
		t.Error("Should return correct answer")
	}
}

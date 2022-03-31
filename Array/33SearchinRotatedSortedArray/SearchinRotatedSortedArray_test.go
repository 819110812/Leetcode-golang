package _3SearchinRotatedSortedArray

import "testing"

func Test_Should_Return_Four(t *testing.T) {
	res := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	if res != 4 {
		t.Error("Expected 4, got ", res)
	}
}

package _210CountHillsandValleysinAnArray

import "testing"

func Test_Should_Return_Three(t *testing.T) {
	res := countHillValley([]int{2, 4, 1, 1, 6, 5})
	if res != 3 {
		t.Errorf("Expected 3, but got %d", res)
	}
}

func Test_Should_Return_Zero(t *testing.T) {
	res := countHillValley([]int{6, 6, 5, 5, 4, 1})
	if res != 0 {
		t.Errorf("Expected 0, but got %d", res)
	}
}

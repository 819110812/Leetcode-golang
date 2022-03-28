package _6RemoveDuplicates

import (
	"testing"
)

func Test_Should_Return_Empty_Array(t *testing.T) {
	res := removeDuplicates([]int{})
	if res != 0 {
		t.Errorf("Expected 0, got %d", res)
	}
}

func Test_Should_Return_OneAndTwo(t *testing.T) {
	// given
	res := removeDuplicates([]int{1, 1, 2})
	if res != 2 {
		t.Errorf("Expected 2,got %v", res)
	}
}

func Test_Should_Return_Five(t *testing.T) {
	res := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	if res != 5 {
		t.Errorf("Expected 5, got %d", res)
	}
}

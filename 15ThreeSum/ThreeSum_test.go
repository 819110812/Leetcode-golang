package _5ThreeSum

import (
	"reflect"
	"sort"
	"testing"
)

func Test_Should_Return_Empty_List(t *testing.T) {
	res := ThreeSum([]int{})
	if len(res) != 0 {
		t.Errorf("Expected empty array, but got %v", res)
	}
	res = ThreeSum([]int{0})
	if len(res) != 0 {
		t.Errorf("Expected empty array, but got %v", res)
	}
}

func Test_Should_Get_Correct_Answer(t *testing.T) {
	testCases := [][]int{
		{-1, 0, 1, 2, -1, -4},
	}
	expected := [][][]int{
		{
			{-1, -1, 2}, {-1, 0, 1},
		},
	}

	for i := 0; i < len(testCases); i++ {
		ans := ThreeSum(testCases[i])
		if !reflect.DeepEqual(ans, expected[i]) {
			t.Errorf("Expected %v, got %v", expected[i], ans)
		}
	}
}

func Test_Should_Not_Contains_Child_Element(t *testing.T) {
	parent := [][]int{
		{-1, -1, 2}, {-1, 0, 1},
	}
	child := []int{-1, 1, 1}
	if Contains(parent, child) {
		t.Errorf("Expected false, but got true")
	}
}

func Test_Should_Sorted(t *testing.T) {
	unstored := []int{3, 2, 1}
	sort.Ints(unstored)
	if !reflect.DeepEqual(unstored, []int{1, 2, 3}) {
		t.Errorf("Expected true, got false")
	}

}

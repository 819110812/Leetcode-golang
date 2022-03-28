package _7RemoveElement

import (
	"testing"
)

func Test_Should_Get_Correct_Answer(t *testing.T) {
	origin := []int{3, 2, 2, 3}
	val := 3
	res := removeElement(origin, val)
	if res != 2 {
		t.Errorf("Expected 2, got %d", res)
	}
}

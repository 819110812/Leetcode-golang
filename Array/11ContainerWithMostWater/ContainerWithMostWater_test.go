package _1ContainerWithMostWater

import (
	"testing"
)

func Test_Should_Return_0(t *testing.T) {
	res := MaxArea([]int{})
	if res != 0 {
		t.Errorf("Expected 0, but got %d", res)
	}
}

func Test_Should_Return_49(t *testing.T) {
	res := MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if res != 49 {
		t.Errorf("Expecte 49， but got %d", res)
	}
}

func Test_Should_Return_1(t *testing.T) {
	res := MaxArea([]int{1, 1})
	if res != 1 {
		t.Errorf("Expected 1, but got %d", res)
	}
}

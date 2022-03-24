package _7LetterCombinationsofaPhoneNumber

import (
	"reflect"
	"testing"
)

func Test_Should_Return_Empty_List(t *testing.T) {
	res := LetterCombinations("")
	if len(res) != 0 {
		t.Errorf("Expected empty string arr, got %v", res)
	}
}

func Test_Should_Return_ShortList(t *testing.T) {
	res := LetterCombinations("2")
	expected := []string{"a", "b", "c"}
	if !isEqual(res, expected) {
		t.Errorf("Expected %v string arr, got %v", expected, res)
	}
}

func Test_Should_Return_LongList(t *testing.T) {
	res := LetterCombinations("23")
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	if !isEqual(res, expected) {
		t.Errorf("Expected %v string arr, got %v", expected, res)
	}
}

func isEqual(arr1 []string, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	if !reflect.DeepEqual(arr1, arr2) {
		return false
	}
	return true
}

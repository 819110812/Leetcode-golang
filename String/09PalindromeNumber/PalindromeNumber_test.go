package PalindromeNumber

import "testing"

func Test_Should_Return_true(t *testing.T) {
	if !IsPalindrome(121) {
		t.Error("Should return true")
	}
	if !IsPalindrome(0) {
		t.Error("Should return true")
	}
}

func Test_Should_Return_false(t *testing.T) {
	if IsPalindrome(-121) {
		t.Error("Should return false")
	}
	if IsPalindrome(123) {
		t.Error("Should return false")
	}
	if IsPalindrome(10) {
		t.Error("Should return false")
	}
}

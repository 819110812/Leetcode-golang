package _4LongestPalindromicSubstring

import "testing"

func Test_ShouldReturnEmpty(t *testing.T) {
	result := LongestPalindromicSubstring("")
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}

func Test_ShouldReturnBAB(t *testing.T) {
	result := LongestPalindromicSubstring("babad")
	if result != "bab" {
		t.Errorf("Expected bab, got %s", result)
	}
}

func Test_ShouldReturnBB(t *testing.T) {
	result := LongestPalindromicSubstring("bb")
	if result != "bb" {
		t.Errorf("Expected bb, got %s", result)
	}
}

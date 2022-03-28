package _3LongestSubsequenceWithoutRepeating

import "testing"

func Test_ShouldReturn_Zero(t *testing.T) {
	lens := LongestSubsequenceWithoutRepeating("")
	if lens != 0 {
		t.Errorf("Expected 0, got %d", lens)
	}
}

func Test_ShouldReturnLengthOne(t *testing.T) {
	lens := LongestSubsequenceWithoutRepeating("a")
	if lens != 1 {
		t.Errorf("expected 1, got %d", lens)
	}
}

func Test_ShouldReturnLengthThree(t *testing.T) {
	lens := LongestSubsequenceWithoutRepeating("abcabcabcabc")
	if lens != 3 {
		t.Errorf("expected 3, got %d", lens)
	}
}

func Test_ShouldAlsoReturnLengthThree(t *testing.T) {
	lens := LongestSubsequenceWithoutRepeating("pwwkew")
	if lens != 3 {
		t.Errorf("expected 3, got %d", lens)
	}
}

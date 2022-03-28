package _4LongestCommonPrefix

import "testing"

func Test_Should_Return_Empty_String(t *testing.T) {
	pre := LongestCommonPrefix([]string{""})
	if pre != "" {
		t.Errorf("Expected emtpy string, got %s", pre)
	}

}

func Test_Should_Return_FL(t *testing.T) {
	pre := LongestCommonPrefix([]string{"flower", "flow", "flight"})
	if pre != "fl" {
		t.Errorf("Expected fl, got %s", pre)
	}
}

func Test_Should_also_Return_Empty(t *testing.T) {
	pre := LongestCommonPrefix([]string{"dog", "racecar", "car"})
	if pre != "" {
		t.Errorf("Expected empty, got %s", pre)
	}
}

func Test_Should_Return_Raw(t *testing.T) {
	pre := LongestCommonPrefix([]string{"a"})
	if pre != "a" {
		t.Errorf("Expected a, got %s", pre)
	}
}

func Test_Should_Return_C(t *testing.T) {
	pre := LongestCommonPrefix([]string{"cir", "car"})
	if pre != "c" {
		t.Errorf("Expected c, got %s", pre)
	}
}

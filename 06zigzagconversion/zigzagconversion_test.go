package _6zigzagconversion

import "testing"

func Test_Return_A(t *testing.T) {
	result := convert("A", 1)
	if result != "A" {
		t.Errorf("Expected A, got %s", result)
	}
}

func Test_Should_Return_Empty(t *testing.T) {
	result := convert("", 0)
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}

func Test_Return_PAYPALISHIRING(t *testing.T) {
	result := convert("PAYPALISHIRING", 3)
	if result != "PAHNAPLSIIGYIR" {
		t.Errorf("Expected PAHNAPLSIIGYIR, got %s", result)
	}
}

func Test_Return_PINALSIGYAHRPI(t *testing.T) {
	result := convert("PAYPALISHIRING", 4)
	if result != "PINALSIGYAHRPI" {
		t.Errorf("Expected PINALSIGYAHRPI, got %s", result)
	}
}

package _7ReverseInteger

import "testing"

func Test_Should_Return_321(t *testing.T) {
	result := Reverse(123)
	if result != 321 {
		t.Error("Expected 321, got ", result)
	}
}

func Test_Should_Return_0(t *testing.T) {
	result := Reverse(0)
	if result != 0 {
		t.Error("Expected 0, got ", result)
	}
}

func Test_Should_Return_Minus321(t *testing.T) {
	result := Reverse(-123)
	if result != -321 {
		t.Error("Expected -321, got ", result)
	}
}

func Test_Should_Return_21(t *testing.T) {
	result := Reverse(120)
	if result != 21 {
		t.Error("Expected 21, got ", result)
	}
}

func Test_should_return_1(t *testing.T) {
	result := Reverse(10)
	if result != 1 {
		t.Error("Expected 1, got ", result)
	}
}

func Test_Should_Return_negative1(t *testing.T) {
	result := Reverse(-10)
	if result != -1 {
		t.Error("Expected -1, got ", result)
	}
	result = Reverse(-1)
	if result != -1 {
		t.Error("Expected -1, got ", result)
	}
}

func Test_Should_Return_negative9(t *testing.T) {
	result := Reverse(-900000)
	if result != -9 {
		t.Error("Expected -9, got ", result)
	}
}

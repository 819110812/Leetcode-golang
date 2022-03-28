package _2IntegertoRoman

import (
	"fmt"
	"testing"
)

func Test_Should_Return_Empty_String(t *testing.T) {
	// Given
	input := 0

	// When
	result := IntToRoman(input)

	// Then
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}

func Test_Should_Return_I(t *testing.T) {
	// given
	input := 1

	// when
	result := IntToRoman(input)

	// then
	if result != "I" {
		t.Errorf("Expected I, got %s", result)
	}
}

func Test_Should_Return_LVIII(t *testing.T) {
	// given
	input := 58

	// when
	result := IntToRoman(input)

	if result != "LVIII" {
		t.Errorf("Expected LVIII, got %s", result)
	}
}

func Test_Should_Return_III(t *testing.T) {
	// given
	input := 3

	// when
	result := IntToRoman(input)

	if result != "III" {
		t.Errorf("Expected III, got %s", result)
	}
}

func Test_Should_Return_MCMXCIV(t *testing.T) {
	// given
	input := 1994

	// when
	result := IntToRoman(input)

	// then
	if result != "MCMXCIV" {
		t.Errorf("Expected MCMXCIV, got %s", result)
	}

}

func Test_Return_Map(t *testing.T) {
	romans := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	for k, v := range romans {
		fmt.Printf("key is %d,value is %s", k, v)
	}
}

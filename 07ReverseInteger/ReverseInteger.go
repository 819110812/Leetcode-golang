package _7ReverseInteger

import "fmt"

func Reverse(x int) int {
	var result int
	for x >= 10 || x < -1 {
		result = result*10 + x%10
		x /= 10
	}
	fmt.Println(x)
	if x != 0 {
		result = result*10 + x
	}
	if result > 2147483647 || result < -2147483648 {
		return 0
	}
	return result
}

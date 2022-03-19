package PalindromeNumber

func IsPalindrome(x int) bool {
	temp := x
	if x == 0 {
		return true
	}
	if x < 0 || (x%10 == 0) {
		return false
	}
	var reverse int
	for x >= 10 {
		reverse = x%10 + reverse*10
		x = x / 10
	}
	if x > 0 {
		reverse = x%10 + reverse*10
	}
	return reverse == temp

}

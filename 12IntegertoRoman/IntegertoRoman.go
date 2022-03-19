package _2IntegertoRoman

func IntToRoman(num int) string {
	// 不能直接遍历map，golang中map是随机无序的
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
	indexes := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	// 如果已经存在了，则直接返回
	if ok := romans[num]; len(ok) > 0 {
		return ok
	}
	res := make([]byte, 0)

	for _, value := range indexes {
		for num >= value {
			num = num - value
			res = append(res, romans[value]...)
		}
	}

	return string(res)
}

package _7LetterCombinationsofaPhoneNumber

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitsToCharMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	store := make([]string, 0)

	for _, item := range digits {
		list := digitsToCharMap[string(item)]
		if len(store) == 0 {
			store = list
		} else {
			temp := make([]string, 0)
			for _, pre := range store {
				for _, cur := range list {
					str := pre + cur
					temp = append(temp, str)
				}
			}
			store = temp
		}
	}

	return store
}

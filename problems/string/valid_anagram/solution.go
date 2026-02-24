package valid_anagram

func ValidAnagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	aMap := make(map[rune]int)

	for _, ch := range str1 {
		if _, ok := aMap[ch]; ok {
			aMap[ch]++
		} else {
			aMap[ch] = 1
		}
	}

	for _, ch := range str2 {
		if _, ok := aMap[ch]; ok {
			aMap[ch]--
		} else {
			return false
		}
	}

	for _, v := range aMap {
		if v != 0 {
			return false
		}
	}
	return true
}

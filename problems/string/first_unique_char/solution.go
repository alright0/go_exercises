package first_unique_char

func FirstUniqueChar(s string) rune {
	m := make(map[rune]int, len(s))

	for _, ch := range s {
		m[ch]++
	}

	for _, letter := range s {
		if m[letter] == 1 {
			return letter
		}
	}

	return 0
}

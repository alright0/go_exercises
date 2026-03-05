package char_frequency

func CharFrequency(s string) map[rune]int {
	m := map[rune]int{}

	for _, ch := range s {
		m[ch]++
	}

	return m
}

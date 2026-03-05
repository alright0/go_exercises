package longest_substring

func LongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	var maxLen int
	wordMap := make(map[byte]struct{})

	left := 0
	for right := 0; right < len(s); right++ {
		curLet := s[right]
		for {
			if _, ok := wordMap[curLet]; !ok {
				break
			}
			delete(wordMap, s[left])
			left++
		}

		wordMap[curLet] = struct{}{}
		windowLen := right - left + 1
		if maxLen < windowLen {
			maxLen = windowLen
		}

	}
	return maxLen
}

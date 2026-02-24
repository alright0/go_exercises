package group_anagrams

import (
	"testing"
)

type groupAnagramsTestCases struct {
	str    []string
	result int
}

func TestLongestSubstring(t *testing.T) {
	cases := []groupAnagramsTestCases{
		{str: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, result: 3},
		{str: []string{}, result: 0},
	}

	for _, tt := range cases {
		result := GroupAnagrams(tt.str)

		if len(result) != tt.result {
			t.Errorf("GroupAnagrams FAILED: expected len %v got %v", tt.result, result)

		}
	}
}

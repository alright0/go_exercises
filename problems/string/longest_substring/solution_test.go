package longest_substring

import (
	"testing"
)

type longestSubstringTestCases struct {
	str    string
	result int
}

func TestLongestSubstring(t *testing.T) {
	cases := []longestSubstringTestCases{
		{str: "abcabc", result: 3},
		{str: "asdc", result: 4},
		{str: "", result: 0},
	}

	for _, tt := range cases {
		result := LongestSubstring(tt.str)

		if result != tt.result {
			t.Errorf("CharFrequency FAILED: expected %d got %d", tt.result, result)
		}

	}
}

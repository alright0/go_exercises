package first_unique_char

import (
	"testing"
)

type firstUniqueCharTestCases struct {
	str    string
	result rune
}

func TestTreeHeight(t *testing.T) {
	cases := []firstUniqueCharTestCases{
		{str: "abcabddabce", result: 'e'},
		{str: "", result: 0},
	}

	for _, tt := range cases {
		result := FirstUniqueChar(tt.str)

		if result != tt.result {
			t.Errorf("CharFrequency FAILED: expected %d got %d", tt.result, result)
		}

	}
}

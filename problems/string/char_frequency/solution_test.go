package char_frequency

import (
	"testing"
)

type charFrequencyTestCases struct {
	str    string
	result map[rune]int
}

func TestTreeHeight(t *testing.T) {
	cases := []charFrequencyTestCases{
		{str: "abcabddabce", result: map[rune]int{'a': 3, 'b': 3, 'c': 2, 'd': 2, 'e': 1}},
		{str: "", result: map[rune]int{}},
	}

	for _, tt := range cases {
		result := CharFrequency(tt.str)

		if len(result) != len(tt.result) {
			t.Errorf("CharFrequency FAILED: expected map length %d got %d", len(tt.result), len(result))
		}
		for k, v := range result {
			if v != tt.result[k] {
				t.Errorf("CharFrequency FAILED: expected %d got %d", v, tt.result[k])
			}
		}

	}
}

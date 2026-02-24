package valid_anagram

import "testing"

type ValidAnagramTestCase struct {
	Str1   string
	Str2   string
	Result bool
}

func TestValidAnagram(t *testing.T) {
	cases := []ValidAnagramTestCase{
		{Str1: "ADADALDSLV", Str2: "VADADALDSL", Result: true},
		{Str1: "", Str2: "ABF", Result: false},
		{Str1: "YASA", Str2: "", Result: false},
		{Str1: "ASD", Str2: "DFG", Result: false},
	}

	for _, tt := range cases {
		result := ValidAnagram(tt.Str1, tt.Str2)
		if result != tt.Result {
			t.Errorf(`str1: %s str2: %s. fact: %t; I2: expected: %t`, tt.Str1, tt.Str2, result, tt.Result)
		}

	}
}

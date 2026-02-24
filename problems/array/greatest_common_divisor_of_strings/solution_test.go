package greatest_common_divisor_of_strings

import "testing"

type GreatestCommonDivisorOfStringsTestCase struct {
	str1   string
	str2   string
	result string
}

func TestGreatestCommonDivisorOfStrings(t *testing.T) {
	cases := []GreatestCommonDivisorOfStringsTestCase{
		{str1: "ABCABC", str2: "ABC", result: "ABC"},
		{str1: "ABABAB", str2: "ABAB", result: "AB"},
		{str1: "LEET", str2: "CODE", result: ""},
		{str1: "AAAAAB", str2: "AAA", result: ""},
	}

	for _, tt := range cases {
		result := GreatestCommonDivisorOfStrings(tt.str1, tt.str2)
		if result != tt.result {
			t.Errorf(`GreatestCommonDivisorOfStrings FAILED! %s = %s, (%s, %s)`, result, tt.result, tt.str1, tt.str2)
		}
	}
}

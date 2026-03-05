package valid_parentheses

import (
	"testing"
)

type validParenthesesTestCases struct {
	str    string
	result bool
}

func TestValidParentheses(t *testing.T) {
	cases := []validParenthesesTestCases{
		{str: "{{}}[]({})", result: true},
		{str: "{}}}{", result: false},
		{str: "", result: true},
	}

	for _, tt := range cases {
		result := ValidParentheses(tt.str)

		if result != tt.result {
			t.Errorf("CharFrequency FAILED: expected %t got %t", tt.result, result)
		}

	}
}

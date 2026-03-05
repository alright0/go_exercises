package happy_number

import (
	"testing"
)

type isHappyTestCases struct {
	num    int
	result bool
}

func TestIsHappy(t *testing.T) {
	cases := []isHappyTestCases{
		{num: 19, result: true},
		{num: 2, result: false},
	}

	for _, tt := range cases {
		result := IsHappy(tt.num)

		if result != tt.result {
			t.Errorf("IsHappy FAILED: expected %t got %t for num %d", tt.result, result, tt.num)
		}
	}
}

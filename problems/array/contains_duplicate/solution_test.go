package contains_duplicate

import (
	"testing"
)

type ContainsDuplicateTestCase struct {
	Nums   []int
	Result bool
}

func TestContainsDuplicate(t *testing.T) {
	cases := []ContainsDuplicateTestCase{
		{Nums: []int{2, 2, 7, 1}, Result: true},
		{Nums: []int{2, 2, 2, 2}, Result: true},
		{Nums: []int{3}, Result: false},
		{Nums: []int{}, Result: false},
		{Nums: []int{1, 2, 3}, Result: false},
	}

	for _, tt := range cases {
		result := ContainsDuplicate(tt.Nums)
		if result != tt.Result {
			t.Errorf(`%d. expected: %t, fact: %t`, tt.Nums, tt.Result, result)
		}

	}
}

package find_first_and_last_position

import (
	"testing"
)

type SearchRangeTestCases struct {
	arr    []int
	target int
	result [2]int
}

func TestSearchRange(t *testing.T) {
	cases := []SearchRangeTestCases{
		{arr: []int{1, 2, 5, 7, 8, 8, 8, 8, 9, 11, 15, 16}, target: 8, result: [2]int{4, 7}},
		{arr: []int{}, target: 1, result: [2]int{-1, -1}},
		{arr: []int{2, 3, 5, 8}, target: 1, result: [2]int{-1, -1}},
	}

	for _, tt := range cases {
		result := SearchRange(tt.arr, tt.target)

		if result != tt.result {
			t.Errorf("SearchRange FAILED: expected %d got %d", tt.result, result)
		}

	}
}

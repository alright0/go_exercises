package search_in_rotated_sorted_array

import (
	"testing"
)

type SearchInRotatedSortedArrayTestCases struct {
	arr    []int
	target int
	result int
}

func TestSearchInRotatedSortedArray(t *testing.T) {
	cases := []SearchInRotatedSortedArrayTestCases{
		{arr: []int{4, 5, 6, 7, 0, 1, 2}, target: 2, result: 6},
		{arr: []int{}, target: 1, result: -1},
		{arr: []int{5, 8, 2, 3}, target: 1, result: -1},
	}

	for _, tt := range cases {
		result := SearchInRotatedSortedArray(tt.arr, tt.target)

		if result != tt.result {
			t.Errorf("SearchInRotatedSortedArray FAILED: expected %d got %d", tt.result, result)
		}

	}
}

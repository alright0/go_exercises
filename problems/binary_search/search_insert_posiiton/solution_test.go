package search_insert_posiiton

import (
	"testing"
)

type SearchInsertTestCases struct {
	arr    []int
	target int
	result int
}

func TestSearchInsertSortedArray(t *testing.T) {
	cases := []SearchInsertTestCases{
		{arr: []int{1, 5, 10, 13, 22, 39, 44, 54, 58, 64, 91}, target: 55, result: 8},
		{arr: []int{}, target: 1, result: 0},
	}

	for _, tt := range cases {
		result := SearchInsert(tt.arr, tt.target)

		if result != tt.result {
			t.Errorf("SearchInRotatedSortedArray FAILED: expected %d got %d", tt.result, result)
		}

	}
}

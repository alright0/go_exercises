package top_k_frequent_elements

import (
	"testing"
)

type topKFrequentElementsTestCases struct {
	nums   []int
	k      int
	result []int
}

func TestValidParentheses(t *testing.T) {
	cases := []topKFrequentElementsTestCases{
		{nums: []int{1, 3, 1, 3, 2, 1, 4}, k: 3, result: []int{1, 3, 2}},
		{nums: []int{1, 1}, k: 3, result: []int{1}},
		{nums: []int{}, k: 1, result: []int{}},
	}

	for _, tt := range cases {
		result := TopKFrequentElements(tt.nums, tt.k)

		if len(result) != len(tt.result) {
			t.Errorf("TopKFrequentElements FAILED. result len %d != %d", len(result), len(tt.result))
		}
		for i := 0; i < len(result); i++ {
			if result[i] != tt.result[i] {
				t.Errorf("CharFrequency FAILED: expected %d got %d", tt.result[i], result[i])
			}
		}

	}
}

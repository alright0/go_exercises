package top_k_frequent_elements

import (
	"reflect"
	"testing"
)

type TopKFrequentTestCase struct {
	input  []int
	k      int
	result []int
}

func TestBstSearchRecursion(t *testing.T) {
	testCases := []TopKFrequentTestCase{
		{[]int{1, 1, 1, 1, 1, 2, 2}, 2, []int{1, 2}},
		{[]int{1, 2, 3, 1, 2, 1, 3, 1, 3, 3}, 2, []int{1, 3}},
		{[]int{}, 3, []int{}},
	}

	for _, tt := range testCases {
		result := TopKFrequentElements(tt.input, tt.k)
		if !reflect.DeepEqual(result, tt.result) {
			t.Errorf("TopKFrequentElements FAILED! %d != %d", result, tt.result)
			return
		}
	}
}

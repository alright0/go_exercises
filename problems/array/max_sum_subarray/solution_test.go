package max_sum_subarray

import "testing"

type MaxSumSubarrayTestCase struct {
	Array  []int
	K      int
	Result int
}

func TestMaxSumSubarray(t *testing.T) {
	cases := []MaxSumSubarrayTestCase{
		{Array: []int{2, 1, 5, 1, 3, 2}, K: 3, Result: 9},
		{Array: []int{2, 1, 5, 1, 3, 2}, K: 1, Result: 5},
		{Array: []int{}, K: 5, Result: 0},
	}

	for _, tt := range cases {
		result := MaxSumSubarray(tt.Array, tt.K)
		if result != tt.Result {
			t.Errorf(`array fact %d != expected %d`, result, tt.Result)
		}
	}
}

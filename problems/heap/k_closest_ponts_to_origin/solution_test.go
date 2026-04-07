package k_closest_ponts_to_origin

import (
	"reflect"
	"testing"
)

type KClosestPointsTestCase struct {
	input  [][]int
	k      int
	result [][]int
}

func TestKClosestPoints(t *testing.T) {
	testCases := []KClosestPointsTestCase{
		{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{-2, 4}, {3, 3}}},
		{[][]int{}, 1, [][]int{}},
		{[][]int{{1, -1}, {-1, 1}, {1, 1}, {-1, -1}}, 2, [][]int{{-1, -1}, {-1, 1}}},
	}

	for _, tt := range testCases {
		result := TopKClosestPointsToOrigin(tt.input, tt.k)
		if !reflect.DeepEqual(result, tt.result) {
			t.Errorf("TopKClosestPointsToOrigin FAILED! %d != %d", result, tt.result)
			return
		}
	}
}

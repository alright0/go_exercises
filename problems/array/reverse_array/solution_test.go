package reverse_array

import "testing"

type ReverseArrayTestCase struct {
	Array  []int
	Result []int
}

func TestReverseArray(t *testing.T) {
	cases := []ReverseArrayTestCase{
		{Array: []int{1, 2, 3, 4, 5}, Result: []int{5, 4, 3, 2, 1}},
		{Array: []int{1}, Result: []int{1}},
		{Array: []int{}, Result: []int{}},
	}

	for _, tt := range cases {
		result := ReverseArray(tt.Array)
		if len(result) != len(tt.Result) {
			t.Errorf(`array fact %d != expected %d`, result, tt.Result)
		}
		for i := range result {
			if result[i] != tt.Result[i] {
				t.Errorf(`array fact %d != expected %d`, result, tt.Result)
			}
		}
	}
}

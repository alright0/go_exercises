package two_sum

import (
	"testing"
)

type TwoSumTestCase struct {
	Nums   []int
	Target int
	I1     int
	I2     int
}

func TestTwoSumBruteforce(t *testing.T) {
	cases := []TwoSumTestCase{
		{Nums: []int{2, 7, 11, 15}, Target: 9, I1: 0, I2: 1},
		{Nums: []int{1, 2, 3, 4}, Target: 5, I1: 1, I2: 2},
		{Nums: []int{}, Target: 10, I1: 0, I2: 0},
		{Nums: []int{11, 15}, Target: 10, I1: 0, I2: 0},
	}

	for _, tt := range cases {
		index1, index2 := TwoSumBruteforce(tt.Nums, tt.Target)
		if index1 != tt.I1 || index2 != tt.I2 {
			t.Errorf(`%d. I1: %d fact: %d; I2: expected: %d fact: %d`, tt.Nums, tt.I1, index1, tt.I2, index2)
		}

	}
}

func TestTwoSumOptimal(t *testing.T) {
	cases := []TwoSumTestCase{
		{Nums: []int{2, 7, 11, 15}, Target: 9, I1: 0, I2: 1},
		{Nums: []int{1, 2, 3, 4}, Target: 5, I1: 1, I2: 2},
		{Nums: []int{}, Target: 10, I1: 0, I2: 0},
		{Nums: []int{11, 15}, Target: 10, I1: 0, I2: 0},
	}

	for _, tt := range cases {
		index1, index2 := TwoSumOptimal(tt.Nums, tt.Target)
		if index1 != tt.I1 || index2 != tt.I2 {
			t.Errorf(`%d. I1: %d fact: %d; I2: expected: %d fact: %d`, tt.Nums, tt.I1, index1, tt.I2, index2)
		}

	}
}

func BenchmarkTwoSumBruteforce(b *testing.B) {
	nums := make([]int, 1000000)

	for i := 0; i < b.N; i++ {
		TwoSumBruteforce(nums, 9999999)
	}
}

func BenchmarkTwoSumOptimal(b *testing.B) {
	nums := make([]int, 1000000)

	for i := 0; i < b.N; i++ {
		TwoSumBruteforce(nums, 9999999)
	}
}

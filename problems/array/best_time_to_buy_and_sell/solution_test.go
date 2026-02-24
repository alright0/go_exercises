package best_time_to_buy_and_sell

import "testing"

type FindByAndSellTestCase struct {
	Prices []int
	Result int
}

func TestFindByAndSell(t *testing.T) {
	cases := []FindByAndSellTestCase{
		{Prices: []int{7, 1, 5, 3, 6, 4}, Result: 5},
		{Prices: []int{}, Result: 0},
		{Prices: []int{1, 2, 3}, Result: 2},
		{Prices: []int{3, 2, 1}, Result: 0},
	}

	for _, tt := range cases {
		result := FindByAndSell(tt.Prices)
		if result != tt.Result {
			t.Errorf(`Prices: %d fact: %d; expected: %d`, tt.Prices, result, tt.Result)
		}

	}
}

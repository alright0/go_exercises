package best_time_to_buy_and_sell_2

import "testing"

type FindByAndSellTwoDealsTestCase struct {
	Prices []int
	Result int
}

func TestFindByAndSellTwoDeals(t *testing.T) {
	cases := []FindByAndSellTwoDealsTestCase{
		{Prices: []int{7, 1, 5, 3, 6, 4}, Result: 7},
		{Prices: []int{}, Result: 0},
		{Prices: []int{1, 2, 3, 1, 2, 3}, Result: 4},
		{Prices: []int{3, 2, 1}, Result: 0},
	}

	for _, tt := range cases {
		result := FindByAndSellTwoDeals(tt.Prices)
		if result != tt.Result {
			t.Errorf(`Prices: %d fact: %d; expected: %d`, tt.Prices, result, tt.Result)
		}

	}
}

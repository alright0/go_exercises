package best_time_to_buy_and_sell_2

import (
	"math"
)

func FindByAndSellTwoDeals(prices []int) int {
	buy1 := math.MaxInt
	sell1 := 0
	buy2 := math.MaxInt
	sell2 := 0

	for _, price := range prices {
		buy1 = min(buy1, price)
		sell1 = max(sell1, price-buy1)

		buy2 = min(buy2, price-sell1)
		sell2 = max(sell2, price-buy2)
	}
	return sell2
}

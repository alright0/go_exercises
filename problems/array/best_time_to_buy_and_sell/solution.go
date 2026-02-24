package best_time_to_buy_and_sell

func FindByAndSell(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var minPrice = prices[0]
	var maxProfit int
	var profit int

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
			continue
		} else {
			profit = price - minPrice
		}
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

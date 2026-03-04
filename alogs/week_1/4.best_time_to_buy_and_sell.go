package main

import "fmt"

func findByAndSell(prices []int) int {
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

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	result := findByAndSell(prices)
	fmt.Println(result)
}

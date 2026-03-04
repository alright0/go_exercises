package main

import (
	"fmt"
	"math"
)

func getProfit(prices []int) int {
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

func main() {
	prices := []int{1, 2, 3, 9, 5, 2, 4, 5}

	result := getProfit(prices)
	fmt.Println(result)
}

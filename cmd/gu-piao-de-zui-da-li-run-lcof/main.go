package main

func maxProfit(prices []int) int {
	minPrice := 2147483647
	theMaxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if x := price - minPrice; x > theMaxProfit {
			theMaxProfit = x
		}
	}
	return theMaxProfit
}

package main

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxprofit := 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if profit := p - minPrice; profit > maxprofit {
			maxprofit = profit
		}
	}
	return maxprofit
}

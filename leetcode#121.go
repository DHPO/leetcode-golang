package main

func maxProfit(prices []int) int {
	maxProfit := 0
	minV := 110000000
	for _, v := range prices {
		if v < minV {
			minV = v
		} else {
			currProfit := v - minV
			if currProfit > maxProfit {
				maxProfit = currProfit
			}
		}
	}
	return maxProfit
}
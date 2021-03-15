package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {
	buy1 := -1000
	sell1 := -1000
	buy2 := -1000
	sell2 := -1000

	for _, p := range prices {
		buy1 = max(buy1, -p)
		sell1 = max(sell1, buy1 + p)
		buy2 = max(buy2, sell1-p)
		sell2 = max(sell2, buy2+p)
	}

	return sell2
}
package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {
	a, b, c := -1000000, 0, 0
	for _, p := range prices {
		a_ := max(a, c-p)
		b_ := max(b, a+p)
		c_ := max(c, b)
		a, b, c = a_, b_, c_
	}
	return b
}
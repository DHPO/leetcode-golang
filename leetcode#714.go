package main

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func maxProfit(prices []int, fee int) int {
	with, without := -1000000, 0
	for _, p := range prices {
		with_ := max(with, without - p - fee)
		without_ := max(without, with + p)
		with, without = with_, without_
	}
	return without
}
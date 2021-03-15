package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(k int, prices []int) int {
	buy := make([]int, k)
	sell := make([]int, k)

	for i := range buy {
		buy[i] = -100000
		sell[i] = -100000
	}

	for _, p := range prices {
		buy[0] = max(buy[0], -p)
		sell[0] = max(sell[0], buy[0]+p)
		for i := 1; i < k; i ++ {
			buy[i] = max(buy[i], sell[i-1]-p)
			sell[i] = max(sell[i], buy[i]+p)
		}
	}

	return sell[k-1]
}
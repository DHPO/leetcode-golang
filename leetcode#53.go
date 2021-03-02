package main

func maxSubArray(nums []int) int {
	PreSum := make([]int, len(nums))
	preSum := 0
	for i, v := range nums {
		preSum += v
		PreSum[i] = preSum
	}
	return maxProfit(PreSum)
}

func maxProfit(prices []int) int {
	maxProfit := -10000000
	minV := 110000000
	for _, v := range prices {
		if v < minV {
			minV = v
		}
		currProfit := v - minV
		if currProfit > maxProfit {
			maxProfit = currProfit
		}
	}
	return maxProfit
}

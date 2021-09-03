package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSubArray(nums []int) int {
	currMax := -1000000
	Max := currMax

	for _, n := range nums {
		currMax = max(currMax + n, n)
		Max = max(currMax, Max)
	}

	return Max
}
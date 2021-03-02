package main

func subarraySum(nums []int, k int) int {
	preSum := 0
	preSumMap := make(map[int]int)
	preSumMap[0] = 1

	result := 0
	for _, v := range(nums) {
		preSum += v
		result += preSumMap[preSum - k]
		preSumMap[preSum] += 1
	}

	return result
}
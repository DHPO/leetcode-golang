package main

func lengthOfLIS(nums []int) int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = 1
	}
	max := 1
	for i, n := range nums {
		for j := 0; j < i; j ++ {
			if nums[j] < n && result[j] + 1 > result[i] {
				result[i] = result[j] + 1
			}
			if result[i] > max {
				max = result[i]
			}
		}
	}
	return max
}
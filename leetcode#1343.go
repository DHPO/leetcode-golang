package main

func numOfSubarrays(arr []int, k int, threshold int) int {
	currSum := 0
	result := 0
	threshold = threshold * k
	for _, v := range arr[:k] {
		currSum += v
	}
	if currSum > threshold {
		result += 1
	}
	for i := k; i < len(arr); i ++ {
		currSum -= arr[i - k]
		currSum += arr[i]
		if currSum > threshold {
			result += 1
		}
	}
	return result
}
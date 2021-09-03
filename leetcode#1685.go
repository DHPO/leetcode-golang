package main

func getSumAbsoluteDifferences(nums []int) []int {
	allSum := 0
	for _, v := range nums {
		allSum += v
	}
	count := len(nums)

	preSum := 0
	postSum := allSum
	result := make([]int, count)
	for i, v := range nums {
		preCount := i
		postCount := count - i
		result[i] = preCount * v - preSum + postSum - postCount * v
		preSum += v
		postSum -= v
	}

	return result
}

func main() {
	getSumAbsoluteDifferences([]int{2, 3, 5})
}
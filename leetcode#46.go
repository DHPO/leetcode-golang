package main

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	lastResult := permute(nums[1:])
	results := [][]int{}

	for i := 0; i < len(nums); i ++{
		for _, r := range lastResult {
			result := make([]int, len(nums))
			copy(result[:i], r[:i])
			copy(result[i + 1:], r[i:])
			result[i] = nums[0]
			results = append(results, result)
		}
	}

	return results
}

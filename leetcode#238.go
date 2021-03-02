package main

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	result[0] = 1
	for i := 1; i < len(nums); i ++ {
		result[i] = nums[i - 1] * result[i - 1]
	}
	rightProduct := nums[len(nums) - 1]
	for i := len(nums) - 2; i >= 0; i -- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	} 

	return result
}
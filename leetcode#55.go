package main

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	curr, max := 1, nums[0]
	for ; curr <= max && max < len(nums); {
		if curr + nums[curr] > max {
			max = curr + nums[curr]
		}
		curr ++
	}
	return max >= len(nums) - 1
}

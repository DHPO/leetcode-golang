package main

func findMin(nums []int) int {
	N := len(nums)
	if nums[0] < nums[N - 1] {
		return nums[0]
	}
	i, j := 0, N
	for {
		mid := (i + j) / 2
		if nums[mid] > nums[j] {
			i = mid + 1
		} else if nums[mid] < nums[j] {
			j = mid
		} else {
			j -= 1
		}
		if i == j {
			return nums[i]
		}
	}
}

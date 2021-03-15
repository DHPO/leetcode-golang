package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	maxLen := 1
	maxLenIdx := 0
	result := make([]int, len(nums))
	before := make([]int, len(nums))

	result[0] = 1
	before[0] = 0

	for i := 1; i < len(nums); i ++ {
		length := 1
		before[i] = i
		for j := i - 1; j >= 0; j -- {
			if nums[i] % nums[j] == 0 && result[j] + 1 > length {
				length = result[j] + 1
				before[i] = j
			}
		}
		result[i] = length
		if length > maxLen {
			maxLen = length
			maxLenIdx = i
		}
	}

	ret := make([]int, 0)
	idx := maxLenIdx
	for {
		ret = append(ret, nums[idx])
		if before[idx] == idx {
			break
		}
		idx = before[idx]
	}

	return ret
}

func main() {
	largestDivisibleSubset([]int{1,2,3})
}
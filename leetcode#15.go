package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	currI := -1000000
	for i := 0; i < len(nums) - 2; i++ {
		for i < len(nums) - 2 && currI == nums[i] {i++}
		currI = nums[i]
		residual := -currI
		if residual < 0 {
			break
		}
		for j, k := i+1, len(nums)-1; j < k; {
			if nums[j] + nums[k] == residual {
				currJ, currK := nums[j], nums[k]
				result = append(result, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == currJ {j++}
				for j < k && nums[k] == currK {k--}
			} else if nums[j] + nums[k] < residual {
				j ++
			} else {
				k --
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}

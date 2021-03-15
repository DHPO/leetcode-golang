package main

import "sort"

func nextPermutation(nums []int)  {
    stack := []int{}
    for i := len(nums) - 1; i >= 0; i -- {
        if len(stack) > 0 && nums[stack[len(stack) - 1]] > nums[i] {
            idx := stack[len(stack) - 1]
            for len(stack) > 0 && nums[stack[len(stack) - 1]] > nums[i]{
                idx = stack[len(stack) - 1]
                stack = stack[: len(stack) - 1]
            }
            nums[i], nums[idx] = nums[idx], nums[i]
			sort.Slice(nums[i + 1:], func (j, k int) bool {return nums[i + 1:][j] < nums[i + 1:][k]})
            return
        } else {
            stack = append(stack, i)
        }
    }
    sort.Slice(nums, func (i, j int) bool {return nums[i] < nums[j]})
}

func main() {
	nextPermutation([]int{1, 3, 2})
}
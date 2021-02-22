package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	result := append([]int{}, nums...)
	length := len(nums)
	for l := 2; l <= length; l++ {
		for s := 0; s <= length - l; s++ {
			sum := (result[s + 1] + nums[s])
			if sum == 0 || (k != 0 && sum % k == 0){
				return true
			} else {
				result[s] = sum
			}
		}
	}
	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{23,6,4,7}, 6))
}
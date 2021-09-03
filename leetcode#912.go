package main

import "fmt"

func move(nums []int) int {
	writeIdx := 1
	pivot := nums[0]
	for i, v := range nums[1:]{
		if v < pivot {
			nums[i + 1], nums[writeIdx] = nums[writeIdx], nums[i + 1]
			writeIdx += 1
		}
	}
	nums[0], nums[writeIdx - 1] = nums[writeIdx - 1], nums[0]
	return writeIdx - 1
}

func sortArray(nums []int) []int {
	if len(nums)  > 1 {
		mid := move(nums)
		sortArray(nums[:mid])
		sortArray(nums[mid+1:])
	}
	return nums
}

func main() {
	data := []int{3,5,1,4,2}
	sortArray(data)
	fmt.Println(data)
}
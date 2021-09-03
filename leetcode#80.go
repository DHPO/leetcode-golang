package main

func removeDuplicates(nums []int) int {
	length := 0
	writeIdx := 0
	currValue, currCount := 0, 0

	for _, v := range nums {
		if v != currValue {
			currValue = v
			currCount = 1
		} else {
			currCount += 1
		}

		if currCount <= 2 {
			nums[writeIdx] = v
			writeIdx += 1
			length += 1
		}
	}

	return length
}
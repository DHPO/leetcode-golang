package main

func findMaxConsecutiveOnes(nums []int) int {
    a, b := 0, 0
	result := 0
	for _, n := range nums {
		if n == 0 {
			b = a + 1
			a = 0
		} else {
			a += 1
			b += 1
		}
		if b > result {
			result = b
		}
	}
	return result
}
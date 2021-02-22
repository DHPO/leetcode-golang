package main

import "fmt"

func trap(height []int) int {
	total := 0
	currMax, currMaxIdx := 0, 0
	currRain := 0
	for i := 0; i < len(height); i ++ {
		if height[i] >= currMax {
			currMax = height[i]
			currMaxIdx = i
			total += currRain
			currRain = 0
		} else {
			currRain += currMax - height[i]
		}
	}
	currRain = 0
	currMax = 0
	for i := len(height) - 1; i >= currMaxIdx; i -- {
		if height[i] >= currMax {
			currMax = height[i]
			total += currRain
			currRain = 0
		} else {
			currRain += currMax - height[i]
		}
	}
	return total
}

func main() {
	fmt.Println(trap([]int{4,2,0,3,2,5}))
}
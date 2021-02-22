package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxValue(grid [][]int) int {
	gridLen := len(grid[0])
	value := make([]int, gridLen)
	value[0] = grid[0][0]
	for i := 1; i < gridLen; i++ {
		value[i] = value[i-1] + grid[0][i]
	}
	for j := 1; j < len(grid); j++ {
		currValue := make([]int, gridLen)
		currValue[0] = grid[j][0] + value[0]
		for i := 1; i < gridLen; i++ {
			currValue[i] = max(currValue[i-1], value[i]) + grid[j][i]
		}
		value = currValue
	}
	return value[gridLen - 1]
}

func main() {
	fmt.Println(maxValue([][]int{{1,3,1},{1,5,1},{4,2,1}}))
}
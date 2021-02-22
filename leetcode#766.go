package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 {
		return true
	}
	cols := len(matrix[0])
	rows := len(matrix)
	for i := 0; i < cols; i++ {
		val := matrix[0][i]
		for j := 1; j < rows && (i + j) < cols; j++ {
			if matrix[j][i + j] != val {
				return false
			}
		}
	}
	for i := 1; i < rows; i++ {
		val := matrix[i][0]
		for j := 0; j < cols && (i + j) < rows; j++ {
			if matrix[i][i + j] != val {
				return false
			}
		}
	}
	return true
}

func main () {
	fmt.Println(isToeplitzMatrix([][]int{{1,2,3,4},{5,1,2,3},{9,5,1,2}}))
}
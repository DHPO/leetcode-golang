package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	imax := len(matrix) - 1
	jmax := len(matrix[0]) - 1
	i, j := 0, jmax

	for {
		if j < 0 || i > imax {
			return false
		}
		if v := matrix[i][j]; v == target {
			return true
		} else if v < target {
			i += 1
		} else {
			j -= 1
		}
	}
}

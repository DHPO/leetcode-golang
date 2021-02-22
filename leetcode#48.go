package main

func rotate(matrix [][]int)  {
	if len(matrix) == 0 {
		return
	}
	n := len(matrix)
	for i := 0; i < int((n + 1) / 2); i ++ {
		for j := 0; j < int(n / 2); j ++ {
			matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i], matrix[i][j] = matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i]
		}
	}
}

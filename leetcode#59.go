package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	minI, minJ, maxI, maxJ := 0, 0, n - 1, n - 1
	i, j := 0, 0
	curr := 1

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIdx := 0

	for {
		dir := dirs[dirIdx]

		deltaI, deltaJ := dir[0], dir[1]

		for i >= minI && i <= maxI && j >= minJ && j <= maxJ {
			matrix[i][j] = curr
			curr += 1
			i += deltaI
			j += deltaJ
		}

		if curr == n * n {
			break
		}

		if dirIdx == 0 {
			minI += 1
		} else if dirIdx == 1 {
			maxJ -= 1
		} else if dirIdx == 2 {
			maxI -= 1
		} else {
			minJ += 1
		}

		dirIdx += 1
	}

	return matrix
}
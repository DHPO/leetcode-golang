package main

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	maxI := len(matrix)
	maxJ := len(matrix[0])
	pathStart := make([][]int, maxI)
	for i := 0; i < maxI; i ++ {
		pathStart[i] = make([]int, maxJ)
		for j, _ := range pathStart[i] {
			pathStart[i][j] = -1
		}
	}

	var dfs func (i, j int) int
	dfs = func (i, j int) int {
		if pathStart[i][j] != -1 {
			return pathStart[i][j]
		}
		maxLen := 1
		if i > 0 && matrix[i - 1][j] > matrix[i][j] && dfs(i - 1, j) >= maxLen {
			maxLen = pathStart[i - 1][j] + 1
		}
		if i < maxI - 1 && matrix[i + 1][j] > matrix[i][j] && dfs(i + 1, j) >= maxLen {
			maxLen = pathStart[i + 1][j] + 1
		}
		if j > 0 && matrix[i][j - 1] > matrix[i][j] && dfs(i, j - 1) >= maxLen {
			maxLen = pathStart[i][j - 1] + 1
		}
		if j < maxJ - 1 && matrix[i][j + 1] > matrix[i][j] && dfs(i, j + 1) >= maxLen {
			maxLen = pathStart[i][j + 1] + 1
		}
		pathStart[i][j] = maxLen
		return maxLen
	}

	result := 0
	for i := 0; i < maxI; i ++ {
		for j := 0; j < maxJ; j ++ {
			if dfs(i, j) > result {
				result = pathStart[i][j]
			}
		}
	}

	return result
}
package main

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func uniquePaths(m int, n int) int {
	result := 1
	for i := 0; i < min(m, n) - 1; i ++ {
		result *= (m + n - 2 - i)
	}
	for i := 0; i < min(m, n) - 1; i ++ {
		result /= i + 1
	}
	return result
}
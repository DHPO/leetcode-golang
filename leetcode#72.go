package main

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1) + 1)
	for i, _ := range dp {
		dp[i] = make([]int, len(word2) + 1)
	}
	for j := 1; j < len(word2) + 1; j ++ {
		dp[0][j] = j
	}
	for i := 1; i < len(word1) + 1; i ++ {
		dp[i][0] = i
		for j := 1; j < len(word2) + 1; j ++ {
			var result int
			if word1[i] == word2[j] {
				result = min(dp[i-1][j-1], min(dp[i-1][j] + 1, dp[i][j-1] + 1))
			} else {
				result = min(dp[i-1][j-1] + 1, min(dp[i-1][j] + 1, dp[i][j-1] + 1))
			}
			dp[i][j] = result
		}
	}
	return dp[len(word1) + 1][len(word2) + 1]
}

package main

func IF(cond bool, a, b int) int {
	if cond {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	return IF(a > b, a, b)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	curr := make([]int, len(text2) + 1)

	for i := 0; i < len(text1); i ++ {
		prev := curr
		curr = make([]int, len(text2) + 1)
		for j := 1; j < len(text2) + 1; j ++ {
			curr[j] = Max(prev[j], curr[j - 1])
			if text1[i] == text2[j - 1] {
				curr[j] = Max(curr[j], prev[j - 1] + 1)
			}
		}
	}

	return curr[len(curr) - 1]
}
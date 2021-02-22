package main

func numTrees(n int) int {
	buf := []int{0: 1, 1: 1}
	for i := 2; i <=n; i ++ {
		result := 0
		for j := 0; j <= i - 1; j++ {
			result += buf[j] * buf[i - 1 - j]
		}
		buf = append(buf, result)
	}
	return buf[n]
}
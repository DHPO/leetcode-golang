package main

func numSquares(n int) int {
	squares := make([]int, 100)
	for i := 0; i < 100; i ++ {
		squares[i] = (i + 1) * (i + 1)
	}
	results := make([]int, n + 1)
	var dfs func (n int) int
	dfs = func(n int) int {
		if n == 1 {
			return 1
		}
		if results[n] > 0 {
			return results[n]
		}
		result := n
		for i := len(squares) - 1; i >= 0; i -- {
			if squares[i] == n {
				return 1
			} else if squares[i] < n {
				resultI := dfs(n - squares[i]) + 1
				if resultI < result {
					result = resultI
				}
			}
		}
		results[n] = result
		return result
	}
	return dfs(n)
}

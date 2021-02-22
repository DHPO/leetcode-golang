package main

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	a, b := 1, 2
	for i := 2; i < n; i ++ {
		b, a = b + a, b
	}
	return b
}
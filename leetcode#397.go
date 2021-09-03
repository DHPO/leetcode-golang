package main

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func integerReplacement(n int) int {
	mem := make(map[int]int)

	var f func(int) int
	f = func(n int) int {
		if value, ok := mem[n]; ok {
			return value
		}
		if n == 1 {
			return 0
		}
		if n % 2 == 0 {
			return f(n / 2) + 1
		}
		result := min(f(n - 1), f(n + 1)) + 1
		mem[n] = result
		return result
	}

	return f(n)
}
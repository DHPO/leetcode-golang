package main

import "fmt"

const Mod int = 1e9 + 7

var matrix = [...]int{2, 2, 2, 3}

func Multi(a, b [4]int) [4]int {
	return [4]int {
		(a[0] * b[0] + a[1] * b[2]) % Mod,
		(a[0] * b[1] + a[1] * b[3]) % Mod,
		(a[2] * b[0] + a[3] * b[2]) % Mod,
		(a[2] * b[1] + a[3] * b[3]) % Mod,
	}
}

func Power(n int) [4]int {
	bases := [][4]int{matrix}
	for i := 1; 1 << i <= n; i ++ {
		bases = append(bases, Multi(bases[i - 1], bases[i - 1]))
	}
	result := [4]int{1, 0, 0, 1}
	i := 0
	for n > 0 {
		if n & 1 == 1 {
			result = Multi(result, bases[i])
		}
		n >>= 1
		i += 1
	}
	return result
}

func solution(n int) int {
	m := Power(n - 1)
	return 6 * (m[0] % Mod + m[1] % Mod + m[2] % Mod + m[3] % Mod) % Mod
}

func main() {
	fmt.Println(solution(3))
}
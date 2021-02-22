package main

import "fmt"

func maxRotateFunction(A []int) int {
	F := 0
	sum := 0
	for i, v := range A {
		F += i * v
		sum += v
	}
	max := F
	for i := len(A) - 1; i > 0; i -- {
		F = F + sum - A[i] * len(A)
		if F > max {
			max = F
		}
	}
	return max
}

func main() {
	fmt.Println(maxRotateFunction([]int{}))
}
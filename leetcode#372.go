package main

import (
	"fmt"
)

const mod = 1337

func pow10(n int) int {
	p2 := (n * n) % mod
	p4 := (p2 * p2) % mod
	p8 := (p4 * p4) % mod
	return (p8 * p2) % mod
}

func pow(a, b int) int {
	result := 1
	for i:=0; i < b; i ++ {
		result = (result * a) % mod
	}
	return result
}

func superPow(a int, b []int) int {
	result := 1
	base := a % mod
	for i := len(b) - 1; i >=0; i-- {
		result = (result * pow(base, b[i])) % mod
		base = pow10(base)
	}
	return result % mod
}

func main() {
	fmt.Println(superPow(2, []int{3,0,0}))
}
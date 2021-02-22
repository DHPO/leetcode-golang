package main

import (
	"fmt"
	"math"
)

func arrangeCoins(n int) int {
	result := math.Floor((math.Sqrt(float64(8 * n + 1)) - 1) / 2)
	return int(result)
} 

func main() {
	fmt.Println(arrangeCoins(5))
}
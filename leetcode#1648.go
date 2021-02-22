package main

import (
	"fmt"
	"sort"
)

const mod int = 1e9 + 7

func maxProfit(inventory []int, orders int) int {
	inventory = append(inventory, 0)
	sort.Ints(inventory)
	count, profit, idx := 0, 0, 0
	length := len(inventory)
	for ; idx < length - 1; {
		currInv, prevInv := inventory[length-1-idx], inventory[length-2-idx]
		currCount := (currInv - prevInv) * (idx + 1)
		if currCount + count > orders {
			break
		}
		soldValue := currCount * (currInv + 1 + prevInv) / 2
		profit = (profit + soldValue) % mod 
		count += currCount
		idx++
	}
	packCount :=  (orders - count) / (idx + 1)
	singleCount := orders - count - packCount * (idx + 1)
	profit += (inventory[length-1-idx] * 2 - packCount + 1) * packCount * (idx + 1) / 2 + (inventory[length-1-idx] - packCount) * singleCount

	return profit % mod
}

func main() {
	fmt.Println(maxProfit([]int{3, 5}, 6))
}
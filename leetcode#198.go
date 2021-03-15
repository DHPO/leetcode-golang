package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	an2, an1 := 0, 0
	for _, n := range nums {
		an := max(an2 + n, an1)
		an2, an1 = an1, an
	}
	return an1
}

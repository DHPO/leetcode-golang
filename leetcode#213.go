package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	an2, an1 := 0, nums[0]
	bn2, bn1 := 0, 0

	for i := 1; i < len(nums) - 1; i ++ {
		n := nums[i]
		an := max(an2 + n, an1)
		bn := max(bn2 + n, bn1)
		an2, an1 = an1, an
		bn2, bn1 = bn1, bn
	}
	n := nums[len(nums) - 1]
	bn := max(bn2 + n, bn1)

	return max(bn, an1)
}

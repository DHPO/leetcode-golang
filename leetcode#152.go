package main

func max(nums... int) int {
	result := nums[0]
	for i := 1; i < len(nums); i ++ {
		if nums[i] > result {
			result = nums[i]
		}
	}
	return result
}

func min(nums... int) int {
	result := nums[0]
	for i := 1; i < len(nums); i ++ {
		if nums[i] < result {
			result = nums[i]
		}
	}
	return result
}

func maxProduct(nums []int) int {
	f := nums[0]
	g := nums[0]
	result := f

	for i := 1; i < len(nums); i ++ {
		v := nums[i]
		f_ := max(f * v, v, g * v)
		g_ := min(f * v, v, g * v)
        f = f_
        g = g_
		if f > result {
			result = f
		}
	}

	return result
}
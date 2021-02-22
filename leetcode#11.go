package main

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	i, j := 0, len(height) - 1
	getArea := func(i, j int) int {
		if j <= i {
			return 0
		}
		return min(height[i], height[j]) * (j - i)
	}
	area := getArea(i, j)
	for i < j {
		currI, currJ := i, j
		if height[i] < height[j] {
			for i < j && height[i] <= height[currI] {i++}
		} else {
			for i < j && height[j] <= height[currJ] {j--}
		}
		currArea := getArea(i, j)
		if currArea > area {
			area = currArea
		}
	}
	return area
}
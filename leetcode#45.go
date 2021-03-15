package main

func jump(nums []int) int {
	result := make([]int, len(nums))
	for i, n := range nums {
		currentMinStep := result[i]
		for j := 1; j <= n && i + j < len(nums); j ++ {
			if result[i + j] == 0 || result[i + j] > currentMinStep + 1 {
				result[i + j] = currentMinStep + 1
			}
		}
		if result[len(result) - 1] > 0 {
			break
		}
	}
	return result[len(result) - 1]
}

func main() {
	jump([]int{2,1,3,3,4})
}
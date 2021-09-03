package main

func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}
	lastCode := grayCode(n - 1)
	length := len(lastCode)
	appendCode := make([]int, length)
	for i, v := range lastCode {
		appendCode[length - 1 - i] = v | 1 << (n - 1)
	}
	return append(lastCode, appendCode...)
}
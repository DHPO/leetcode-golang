package main

func OddNum(n int) int {
	return (n + 1) / 2
}

func EvenNum(n int) int {
	return n / 2 + 1
}

func sumOddLengthSubarrays(arr []int) int {
	result := 0
	length := len(arr)

	for i, v := range arr {
		result += v * (OddNum(i) * OddNum(length - i) + EvenNum(i) * EvenNum(length - i))
	}

	return result
}
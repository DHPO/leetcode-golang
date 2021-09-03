package main

import "strconv"

func addTo(target []int, num int, offset int) {
	idx := len(target) - 1 - offset
	for num > 0 {
		currNum := (target[idx] + num) / 10
		target[idx] = (target[idx] + num) % 10
		num = currNum
		idx -= 1
	}
}

func toString(target []int) string {
	result := ""
	idx := 0
	for ; target[idx] == 0; idx ++ {}
	for ; idx < len(target); idx ++ {
		result += strconv.Itoa(target[idx])
	}
	return result
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	digits1 := []byte(num1)
	digits2 := []byte(num2)

	result := make([]int, len(digits1) + len(digits2) + 2)
	for idx1 := len(digits1) - 1; idx1 >= 0; idx1 -- {
		for idx2 := len(digits2) - 1; idx2 >= 0; idx2 -- {
			offset := len(digits1) +  len(digits2) - 2 - idx1 - idx2
			addTo(result, (int(digits1[idx1] - '0') * int(digits2[idx2] - '0')), offset)
		}
	}

	return toString(result)
}
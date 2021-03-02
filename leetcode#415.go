package main

import "strings"

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	result := make([]string, len(num1) + 1, len(num1) + 1)
	res := byte(0)
	for i := 0; i < len(num1); i ++ {
		idx1 := len(num1) - i - 1
		idx2 := len(num2) - i - 1
		idxR := len(num1) - i

		d1 := num1[idx1] - '0'
		var d2 byte
		if idx2 < 0 {
			d2 = 0
		} else {
			d2 = num2[idx2] - '0'
		}
		dr := d1 + d2 + res
		res = dr / 10
		dr = dr % 10

		result[idxR] = string(dr + '0')
	}
	if res == 0 {
		return strings.Join(result[1:], "")
	} else {
		result[0] = string(res + '0')
		return strings.Join(result, "")
	}
}

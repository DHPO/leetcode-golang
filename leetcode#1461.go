package main

import (
	"fmt"
	"strconv"
)

func hasAllCodes(s string, k int) bool {
	if len(s) < k {
        return false
    }
	dict := make([]bool, 1 << k)
	value, _ := strconv.ParseInt(s[:k], 2, 32)
	dict[value] = true
	for _, v := range []byte(s[k:]) {
		value = (value << 1) % (1 << k) + int64(v - '0')
		dict[value] = true
	}
	for _, v := range dict {
		if !v {
			return false
		}
	}
	return true
}

func main () {
	fmt.Println(hasAllCodes("0000000001011100",4))
}
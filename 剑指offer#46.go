package main

import (
	"fmt"
	"strconv"
)

func translateNumStr(num string) int {
	if len(num) == 0 || len(num) == 1 {
		return 1
	}
	if (num[0] == '1') || (num[0] == '2' && num[1] <= '5') {
		return translateNumStr(num[1:]) + translateNumStr(num[2:])
	} else {
		return translateNumStr(num[1:])
	}
}


func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	return translateNumStr(numStr)
}

func main() {
	fmt.Println(translateNum(12258))
}
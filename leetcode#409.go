package main

import "fmt"

func longestPalindrome(s string) int {
	byteCount := make(map[byte]int)
	hasSingle := false
	count := 0

	for _, b := range []byte(s) {
		byteCount[b] += 1
	}

	for _, c := range byteCount {
		if c % 2 == 1 {
			hasSingle = true
		}
		count += c / 2
	}
	count *= 2
	if hasSingle {
		count += 1
	}
	return count
}

func main () {
	fmt.Println(longestPalindrome("abccccdd"))
}

package main

import (
	"fmt"
	"strings"
)

func isCharater(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
        return true
    }
	s = strings.ToLower(s)
	l, r := 0, len(s) - 1
	for ; l < r; {
		for ; l < len(s) &&!isCharater(s[l]); l ++ {}
		for; r >= 0 && !isCharater(s[r]); r -- {}
		if l >= len(s) || r < 0 {
			break
		}
		if s[l] != s[r] {
			return false
		}
		l ++
		r --
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(".."))
}
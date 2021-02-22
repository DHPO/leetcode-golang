package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	s = strings.Join(strings.Split(s, ""), "#")
	result := s[:1]
	maxL := 0
	for center := 1; center < len(s) - maxL; center ++ {
		l := 1
		for ; center - l >= 0 && center + l < len(s) && s[center-l] == s[center+ l]; l ++ {}
		l --
		if l > maxL || ( l == maxL && l % 2 == 1 && s[center] == '#') {
			maxL = l
			result = s[center-l: center+l+1]
		}
	}
	result = strings.ReplaceAll(result, "#", "")
	return result
}

func main() {
	fmt.Println(longestPalindrome("ac"))
}
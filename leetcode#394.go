package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsChar(c byte) bool {
	return !IsDigit(c) && c != '[' && c != ']'
}

func decodeString(s string) string {
	curr := 0
	var decode func() string 
	decode = func () string {
		char := s[curr]
		if IsDigit(char) {
			numStart := curr
			for ; IsDigit(s[curr]); curr ++ {}
			repeat, err := strconv.Atoi(s[numStart: curr])
			if err != nil {
				panic(err)
			}
			results := []string{}
			if s[curr] != '[' {
				panic(fmt.Sprintf("Expect '[' at %d, but found %s", curr, string(s[curr])))
			}
			curr ++
			for s[curr] != ']' {
				str := decode()
				results = append(results, str)
			}
			curr ++
			return strings.Repeat(strings.Join(results, ""), repeat)
		} else {
			start := curr
			for ; curr < len(s) && IsChar(s[curr]); curr ++ {}
			return s[start: curr]
		}
	}
	results := make([]string, 0)
	for ; curr < len(s); {
		results = append(results, decode())
	}
	return strings.Join(results, "")
}

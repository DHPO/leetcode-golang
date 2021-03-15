package main

import "fmt"

func minWindow(s string, t string) string {
	leftIdx, rightIdx := 0, 0
	found, require := make(map[byte]int), make(map[byte]int)

	for _, c := range []byte(t) {
		require[c] += 1
	}
	check := func () bool {
		for char, count := range require {
			if found[char] < count {
				return false
			}
		}
		return true
	}
	for leftIdx < len(s) {
		if _, ok := require[s[leftIdx]]; ok {
			break
		}
		leftIdx += 1
	}
	if leftIdx >= len(s) {
		return ""
	}
	found[s[leftIdx]] += 1
	rightIdx = leftIdx
	result := ""
	for {
		for rightIdx < len(s) - 1 && !check(){
			rightIdx += 1
			found[s[rightIdx]] += 1
		}
		if !check() {
			break
		}
		for leftIdx < len(s) && check() {
			if len(result) == 0 || len(result) > rightIdx - leftIdx + 1 {
				result = s[leftIdx: rightIdx + 1]
			}
			found[s[leftIdx]] -= 1
			leftIdx += 1
		}
		if leftIdx == len(s) - 1 {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(minWindow("AB", "B"))
}
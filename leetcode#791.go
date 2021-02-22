package main

import "sort"

func customSortString(S string, T string) string {
	charMap := make(map[byte]int)
	for i, v := range []byte(S) {
		charMap[v] = i
	}
	str := []byte(T)
	sort.Slice(str, func(i, j int) bool {return charMap[str[i]] < charMap[str[j]]})
	return string(str)
}
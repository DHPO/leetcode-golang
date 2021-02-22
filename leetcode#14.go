package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	result := []byte{}
	sort.Slice(strs, func(i, j int) bool { return strs[i] < strs[j] })
	str1, str2 := strs[0], strs[len(strs) - 1]
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i] {
			result = append(result, str1[i])
		} else {
			break
		}
	}
	return string(result)
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","float"}))
}

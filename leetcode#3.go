package main

import "fmt"

// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 0 {
//         return 0
//     }
// 	prevMap := make(map[byte]int)
// 	prev := make([]int, len(s))
// 	for i, c := range []byte(s) {
// 		if prevIdx, ok := prevMap[c]; ok {
// 			prev[i] = prevIdx
// 		} else {
// 			prev[i] = -1
// 		}
// 		prevMap[c] = i
// 	}
// 	maxLen := 0
// 	buf := make([]bool, len(s))
// 	for i := range buf {
// 		buf[i] = true
// 	}
// 	for l := 1; l < len(s); l ++ {
// 		found := false
// 		for i := 0; i + l < len(s); i++ {
// 			j := i + l
// 			if prev[j] < i && buf[i] {
// 				buf[i] = true
// 				maxLen = l
// 				found = true
// 			} else {
// 				buf[i] = false
// 			}
// 		}
// 		if !found {
// 			break
// 		}
// 	}
// 	return maxLen + 1
// }

func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
	existsMap := make(map[byte]int)
	maxLen := 0
	for i, j := 0, -1; j < len(s); {
		for j < len(s) - 1 && existsMap[s[j + 1]] == 0 {
			existsMap[s[j + 1]] ++
			j ++
		}
		if j - i + 1> maxLen {maxLen = j - i + 1}
		existsMap[s[i]]--
		i += 1
		if i >= len(s) {break}
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
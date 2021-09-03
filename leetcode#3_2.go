package main

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	charCount := make([]int, 256)
	chars := []byte(s)
	start, end := 0, 0

	result := 1
	charCount[chars[0]] += 1
	for end < len(chars) {
		end += 1
		// move end
		for end < len(s) {
			charCount[chars[end]] += 1
			if (charCount[chars[end]] > 1) {
				break
			}
			end += 1
		}
		// calculate length
		length := end - start
		if length > result {
			result = length
		}
		// quit
		if end >= len(chars) {
			break
		}
		// move start
		dupChar := chars[end]
		for {
			charCount[chars[start]] -= 1
			if chars[start] == dupChar {
				start += 1
				break
			}
			start += 1
		}
	}
	return result
}
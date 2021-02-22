package main

import "strings"

func wordPattern(pattern string, s string) bool {
	dict := make(map[byte]string)
	dict2 := make(map[string]byte)

	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	for i := range pattern {
		word := words[i]
		p := pattern[i]
		_, ok1 := dict[p]
		_, ok2 := dict2[word]
		if ok1 != ok2 {
			return false
		}
		if ok1 && (p != dict2[word] || word != dict[p]) {
			return false
		}
		dict[p] = word
		dict2[word] = p
	}

	return true
}

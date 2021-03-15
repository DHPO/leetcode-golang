package main

import (
	"fmt"
	"strings"
)

func attach(s string, sentence []string) []string {
	for i := 0; i < len(sentence); i ++ {
		sentence[i] = strings.Join([]string{s, sentence[i]}, " ")
	}
	return sentence
}

func wordBreak(s string, wordDict []string) []string {
	var match func(s string) []string
	match = func(s string) []string {
		if len(s) == 0 {
			return []string{""}
		}
		result := []string{}
		for _, w := range wordDict {
			if len(s) < len(w) {
				continue
			}
			if s[: len(w)] == w {
				if len(w) == len(s) {
					result = append(result, w)
				} else {
					result = append(result, attach(w, match(s[len(w):]))...)
				}
			}
		}
		return result
	}
	return match(s)
}
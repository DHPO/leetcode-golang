package main

import (
	"fmt"
	"sort"
)

type stack struct {
	data []int
}

func (s *stack) isEmpty() bool {return len(s.data) == 0}

func (s *stack) Push(e int) {s.data = append(s.data, e)}

func (s *stack) Pop() int {
	result := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]
	return result
}

func longestValidParentheses(s string) int {
	unmatchedLeft := stack{}
	unmatchedRight := []int{}

	for i, c := range []byte(s) {
		if c == '(' {
			unmatchedLeft.Push(i)
		} else {
			if unmatchedLeft.isEmpty() {
				unmatchedRight = append(unmatchedRight, i)
			} else {
				unmatchedLeft.Pop()
			}
		}
	}

	unmatched := append(unmatchedRight, unmatchedLeft.data...)
	unmatched = append(unmatched, -1, len(s))
	sort.Ints(unmatched)

	if len(unmatched) == 0 {
		return len(s)
	}

	maxLen := 0
	for i := 0; i < len(unmatched) - 1; i ++ {
		l := unmatched[i + 1] - unmatched[i] - 1
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}

func main() {
	fmt.Println(longestValidParentheses("(()"))
}
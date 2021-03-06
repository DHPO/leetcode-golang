package main

import (
	"sort"
	"strconv"
	"strings"
)

func compare(a, b string) bool {
	swap := false
	if len(a) > len(b) {
		a, b = b, a
		swap = true
	}
	for i := 0; i < len(a); i ++ {
		if a[i] > b[i] {
			return swap != true
		} else if a[i] < b[i] {
			return swap != false
		}
	}
	if len(a) == len(b) {
		return true
	}
	return swap != compare(a, b[len(a):])
}

func largestNumber(nums []int) string {
	numsString := make([]string, len(nums))
	for i, n := range nums {
		numsString[i] = strconv.Itoa(n)
	}
	sort.Slice(numsString, func(i, j int) bool {
		return compare(numsString[i], numsString[j])
	})
	result := strings.Join(numsString, "")
	i := 0
	for i < len(result) && result[i] == '0' {
		i ++
	}
	if i == len(result) {
		return "0"
	} else {
		return result[i:]
	}
}
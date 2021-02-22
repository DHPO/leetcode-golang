package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	if len(people) == 0 {
		return [][]int{}
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] < people[j][0]
	})
	result := make([][]int, len(people))
	h := people[0][0]
	curr, count := 0, 0
	for _, p := range people {
		if p[0] > h {
			curr = 0
			count = 0
			h = p[0]
			for ; result[curr] != nil; curr ++ {}
		}
		for ; count < p[1]; {
			curr ++
			for ; result[curr] != nil; curr ++ {}
			count ++
		}
		result[curr] = p
	}
	return result
}

func main() {
	fmt.Println(reconstructQueue([][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}))
}
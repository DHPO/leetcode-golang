package main

import "sort"

func merge(intervals [][]int) [][]int {
	result := [][]int{}
	if len(intervals) == 0 {
		return result
	}

	sort.Slice(intervals, func (i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	curr := intervals[0]
	for _, interval := range intervals {
		if interval[0] <= curr[1] { // merge
			if interval[1] > curr[1] {
				curr[1] = interval[1]
			}
		} else { // new
			result = append(result, curr)
			curr = interval
		}
	}
	result = append(result, curr)

	return result
}

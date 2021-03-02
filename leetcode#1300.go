package main

import "sort"


func findBestValue(arr []int, target int) int {
	if len(arr) == 0 {
		return target
	}
	sort.Slice(arr, func (i, j int) bool {
		return arr[i] < arr[j]
	})

	PreSum := 0
	for i, v := range arr {
		remainLen := len(arr) - i
		remain := target - PreSum
		if v * remainLen < remain { // set to max
			PreSum += v
			} else {
			ideal := (target - PreSum) / remainLen
			if (remain - ideal * remainLen) > ((ideal + 1) * remainLen - remain) {
				return ideal + 1
			} else {
				return ideal
			}
			break
		}
	}

	return arr[len(arr) - 1]
}
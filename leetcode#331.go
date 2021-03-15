package main

import "strings"

func isValidSerialization(preorder string) bool {
	strs := strings.Split(preorder, ",")
	slotCount := 2

	if len(strs) == 0 || strs[0] == "#" {
		return true
	}

	for _, s := range strs[1:] {
		if slotCount == 0 {
			return false
		}
		if s == "#" {
			slotCount -= 1
		} else {
			slotCount += 1
		}
	}

	return true
}
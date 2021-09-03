package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	versionList1 := strings.Split(version1, ".")
	versionList2 := strings.Split(version2, ".")

	reverse := 1
	if len(versionList1) > len(versionList2) {
		reverse = -1
		versionList1, versionList2 = versionList2, versionList1
	}

	result := 0
	for i := range versionList2 {
		var v1 int
		if i >= len(versionList1) {
			v1 = 0
		} else {
			v1, _ = strconv.Atoi(versionList1[i])
		}
		v2, _ := strconv.Atoi(versionList2[i])
		if v1 < v2 {
			result = -1
			break
		} else if v1 > v2 {
			result = 1
            break
		}
	}

	return reverse * result
}
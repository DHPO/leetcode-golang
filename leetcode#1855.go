package main

func maxDistance(nums1 []int, nums2 []int) int {
	idx1, idx2 := 0, 0
	result := 0
	for {
		if idx1 >= len(nums1) || idx2 >= len(nums2) {
			break
		}
		if nums1[idx1] <= nums2[idx2] {
			length := idx2 - idx1
			if length > result {
				result = length
			}
			idx2 += 1
		} else {
			idx1 += 1
		}
	}

	return result
}
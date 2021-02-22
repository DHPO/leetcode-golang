package main

func min(data... int) int {
	result := 1000000
	for _, v := range data {
		if v < result {
			result = v
		}
	}
	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	start1, start2 := 0, 0
	K, isEven := (len(nums1) + len(nums2) + 1) / 2, (len(nums1) + len(nums2)) % 2 == 0
	var findKth func(int) int
	findKth = func (k int) int {
		if start1 == len(nums1) {
			start2 += k
			return nums2[start2 - 1]
		}
		if start2 == len(nums2) {
			start1 += k
			return nums1[start1 - 1]
		}
		if k == 1 {
			if nums1[start1] < nums2[start2] {
				start1 += 1
				return nums1[start1 - 1]
			} else {
				start2 += 1
				return nums2[start2 - 1]
			}
		}
		mid := min(k / 2, len(nums1) - start1, len(nums2) - start2)
		if nums1[start1 + mid - 1] < nums2[start2 + mid - 1] {
			start1 = start1 + mid
		} else {
			start2 = start2 + mid
		}
		return findKth(k - mid)
	}
	result := findKth(K)
	if isEven {
		result += findKth(1)
		return float64(result) / 2
	} else {
		return float64(result)
	}
}
package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if target < nums[0] && target > nums[len(nums)-1] {
		return -1
	}
	i, j := 0, len(nums) - 1
	if target < nums[0] { // find in latter part
		last := nums[j]
		for i < j - 1 {
			mid := (i + j) / 2
			if nums[mid] == target {
				return mid
			}
			if nums[mid] > last || nums[mid] < target {
				i = mid
			} else {
				j = mid
			}
		}
	} else if target > nums[j]{ // find in former part
		first := nums[j]
		for i < j - 1 {
			mid := (i + j) / 2
			if nums[mid] == target {
				return mid
			}
			if nums[mid] < first || nums[mid] > target {
				j = mid
			} else {
				i = mid
			}
		}
	} else {
        for i < j - 1 {
			mid := (i + j) / 2
			if nums[mid] == target {
				return mid
			}
			if nums[mid] > target {
				j = mid
			} else {
				i = mid
			}
		}
    }
	if nums[j] == target {
		return j
	} else if nums[i] == target {
        return i
    } else {
		return -1
	}
}
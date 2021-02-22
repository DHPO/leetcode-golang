package main

func searchRange(nums []int, target int) []int {
	middle := -1
	// find one
	for i, j := 0, len(nums); i < j; {
		mid := (i + j) / 2
        if i == mid {
            if nums[mid] != target {
                return []int{-1, -1}
            }
        }
		if nums[mid] == target {
			middle = mid
			break
		} else if nums[mid] < target {
			i = mid
		} else {
			j = mid
		}
	}
    first, last := middle, middle
	// find first
	for i, j := 0, middle; i < j; {
		mid := (i + j) / 2
		if i == mid {
			if nums[i] == target {
				first = i
			} else {
				first = i + 1
			}
			break
		}
		if nums[mid] < target {
			i = mid
		} else {
			j = mid
		}
	}
	// find last
	for i, j := middle, len(nums) - 1; i < j; {
		mid := (i + j) / 2
		if i == mid {
			if nums[i + 1] == target {
				last = i + 1
			} else {
				last = i
			}
			break
		}
		if nums[mid] <= target {
			i = mid
		} else {
			j = mid
		}
	}

	return []int{first, last}
}

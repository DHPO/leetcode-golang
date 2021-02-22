package main

func sortColors(nums []int)  {
	r, w, b := -1, -1, -1
	for _, c := range nums {
		if c == 0 {
			r ++
			w ++
			b ++
			nums[b] = 2
			nums[w] = 1
			nums[r] = 0
		} else if c == 1 {
			w ++
			b ++
			nums[b] = 2
			nums[w] = 1
		} else {
			b ++
			nums[b] = 2
		}
	}
}
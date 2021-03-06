package main

import "container/heap"

type Heap []int

func (h Heap) Len() int {return len(h)}
func (h Heap) Less(i, j int) bool {return h[i] > h[j]}
func (h Heap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *Heap) Push(e interface{}) {
	*h = append(*h, e.(int))
}
func (h *Heap) Pop() interface{} {
	result := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]
	return result
}

func maxSlidingWindow(nums []int, k int) []int {
	var h Heap = make([]int, k)
	copy(h, nums[: k])
	heap.Init(&h)
	idxMap := make(map[int][]int)

	result := make([]int, len(nums) - k + 1)
	for i := 0; i < k; i ++ {
		if list, ok := idxMap[nums[i]]; ok {
			idxMap[nums[i]] = append(list, )
		}
	}

	for i := 0; i + k - 1 < len(nums); i ++ {

	}

	return result
}
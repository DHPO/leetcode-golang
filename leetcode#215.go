package main

import "container/heap"

type Heap []int

func (h Heap) Len() int { return len(h)}
func (h Heap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h Heap) Less(i, j int) bool {return h[i] < h[j]}
func (h *Heap) Push(e interface{}) {
	*h = append(*h, e.(int))
}
func (h *Heap) Pop() interface{} {
	result := (*h)[len(*h) - 1]
	*h = (*h)[: len(*h) - 1]

	return result
}

func findKthLargest(nums []int, k int) int {
	h := &Heap{}
	heap.Init(h)

	for _, n := range nums {
		if h.Len() >= k {
			value := (*h)[0]
			if n < value {
				continue
			}
			heap.Pop(h)
		}
		heap.Push(h, n)
	}

	return heap.Pop(h).(int)
}

func main() {
	findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4)
}
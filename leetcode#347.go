package main

import "container/heap"

type Item struct {
	Value int
	Count int
	Index int
}

type Heap []*Item

func (h Heap) Len() int {return len(h)}
func (h Heap) Less(i, j int) bool {return h[i].Count > h[j].Count}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}
func (h *Heap) Push(e interface{}) {
	item := e.(*Item)
	item.Index = len(*h)
	*h = append(*h, item)
}
func (h *Heap) Pop() (result interface{}) {
	result = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return
}


func topKFrequent(nums []int, k int) []int {
	h := &Heap{}
	heap.Init(h)
	m := make(map[int]*Item)

	for _, n := range nums {
		if item, ok := m[n]; ok {
			item.Count += 1
			heap.Fix(h, item.Index)
		} else {
			item := &Item{
				Value: n,
				Count: 1,
			}
			heap.Push(h, item)
			m[n] = item
		}
	}

	result := make([]int, k)
	for i := 0; i < k; i ++ {
		result[i] = heap.Pop(h).(*Item).Value
	}
	return result
}

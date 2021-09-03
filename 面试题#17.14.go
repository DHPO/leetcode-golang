package main

import (
	"container/heap"
)

type Heap struct {
	Values []int
}

func (h *Heap) Len() int {return len(h.Values)}

func (h *Heap) Less(i, j int) bool {return h.Values[i] > h.Values[j]}

func (h *Heap) Swap(i, j int) {h.Values[i], h.Values[j] = h.Values[j], h.Values[i]}

func (h *Heap) Top() int {
	return h.Values[0]
}

func (h *Heap) Push(v interface{}) {
	h.Values = append(h.Values, v.(int))
}

func (h *Heap) Pop() interface{} {
	result := h.Values[len(h.Values) - 1]
	h.Values = h.Values[: len(h.Values) - 1]
	return result
}

func smallestK(arr []int, k int) []int {
	h := &Heap{Values: []int{}}
	heap.Init(h)

	for _, v := range arr {
		if h.Len() < k {
			heap.Push(h, v)
		} else {
			max := h.Top()
			if max > v {
				heap.Pop(h)
				heap.Push(h, v)
			}
		}
	}

	return h.Values
}
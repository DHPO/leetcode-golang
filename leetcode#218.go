package main

import (
	"container/heap"
	"sort"
)

type Building struct {
	height int
	end int
}

type Heap []Building

func (h *Heap) Len() int {return len(*h)}
func (h *Heap) Swap(i, j int) {(*h)[i], (*h)[j] = (*h)[j], (*h)[i]}
func (h *Heap) Less(i, j int) bool { 
	if (*h)[i].height == (*h)[j].height {
		return (*h)[i].end > (*h)[j].end
	}
	return (*h)[i].height > (*h)[j].height
}
func (h *Heap) Push(b interface{}) { *h = append(*h, b.(Building))}
func (h *Heap) Pop() interface{} {
	b := (*h)[h.Len() - 1]
	*h = (*h)[: h.Len() - 1]
	return b
}

func getSkyline(buildings [][]int) [][]int {
	endPos := 0
	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i][0] < buildings[j][0] || (buildings[i][0] == buildings[j][0] && buildings[i][2] > buildings[j][2])
	})
	h := &Heap{}
	heap.Init(h)
	result := [][]int{}

	for _, b := range buildings {
		start, end, height := b[0], b[1], b[2]
		for h.Len() > 0 && (*h)[0].end < start {
			popB := heap.Pop(h).(Building)
			if popB.end < endPos {
				continue
			}
			endPos = popB.end
			for h.Len() > 0 && (*h)[0].end <= endPos {heap.Pop(h)}
			if h.Len() == 0 {
				result = append(result, []int{endPos, 0})
			} else {
				result = append(result, []int{endPos, (*h)[0].height})
			}
		}
		if h.Len() == 0 || (*h)[0].height < height {
			result = append(result, []int{start, height})
		}
		heap.Push(h, Building{height, end})
	}
	for h.Len() > 0 {
		popB := heap.Pop(h).(Building)
		if popB.end < endPos {
			continue
		}
		endPos = popB.end
		for h.Len() > 0 && (*h)[0].end <= endPos {heap.Pop(h)}
		if h.Len() == 0 {
			result = append(result, []int{endPos, 0})
		} else {
			result = append(result, []int{endPos, (*h)[0].height})
		}
	}

	return result
}

func main() {
	getSkyline([][]int{
		{2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8},
	})
}
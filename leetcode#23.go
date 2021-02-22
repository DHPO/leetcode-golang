package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

type item []*ListNode

func (it item) Len() int { return len(it)}

func (it item) Less(i, j int) bool {
	return it[i].Val < it[j].Val
}

func (it item) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func (it *item) Pop() interface{} {
	defer func() {*it = (*it)[: len(*it) - 1]}()
	return (*it)[len(*it) - 1]
}

func (it *item) Push(i interface{}) {
	*it = append(*it, i.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := item{}
	for _, l := range lists {
		if l != nil {
			h = append(h, l)
		}
	}
	heap.Init(&h)

	result := &ListNode{Val: 0}
	curr := result
	for ; h.Len() > 0; {
		l := heap.Pop(&h).(*ListNode)
		curr.Next = l
		curr = curr.Next
		if l.Next != nil {
			heap.Push(&h, l.Next)
		}
	}
	return result.Next
}

func main() {
	fmt.Println()
}
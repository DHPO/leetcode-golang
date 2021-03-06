package main

type Heap []int

func (h *Heap) Init() {
	*h = make([]int, 0)
}

func (h *Heap) Push(e int) {
	*h = append(*h, e)
}

func (h *Heap) Top() int {
	return (*h)[len(*h) - 1]
}

func (h *Heap) Pop() int {
	result := h.Top()
	*h = (*h)[: len(*h) - 1]
	return result
}

func (h *Heap) IsEmpty() bool {
	return len(*h) == 0
}

func validateStackSequences(pushed []int, popped []int) bool {
	pushIdx, popIdx := 1, 0
	h := Heap{}
	h.Init()

	h.Push(pushed[0])
	for pushIdx < len(pushed) {
		for !h.IsEmpty() && h.Top() == popped[popIdx] {
			h.Pop()
			popIdx ++
		}
		h.Push(pushed[pushIdx])
		pushIdx ++
	}
	for !h.IsEmpty() && h.Top() == popped[popIdx] {
		h.Pop()
		popIdx ++
	}
	return h.IsEmpty()
}
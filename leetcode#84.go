package main

type Stack struct {
	data []int
}

func (this *Stack) Push(e int) {
	this.data = append(this.data, e)
}

func (this *Stack) Pop() (e int) {
	if len(this.data) == 0 {
		return 0
	}
	e = this.data[len(this.data) - 1]
	this.data = this.data[:len(this.data)-1]
	return
}

func (this *Stack) Top() (e int) {
	if len(this.data) == 0 {
		return 0
	}
	e = this.data[len(this.data) - 1]
	return
}

func (this *Stack) IsEmpty() bool {
	return len(this.data) == 0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := Stack{}
	result := 0

	for i, v := range heights {
		for !stack.IsEmpty() && heights[stack.Top()] >= v {
			h := heights[stack.Pop()]
			var idx int
			if stack.IsEmpty() {
				idx = -1
			} else {
				idx = stack.Top()
			}
			w := (i - idx - 1)
			result = max(result, h * w)
		}
		stack.Push(i)
	}

	return result
}

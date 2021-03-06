package main

type Stack struct {
	data []int
}

func (s *Stack) Init() {
	s.data = make([]int, 0)
}

func (s *Stack) Top() int {
	return s.data[len(s.data) - 1]
}

func (s *Stack) Pop() int {
	result := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]
	return result
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(e int) {
	s.data = append(s.data, e)
}

func nextGreaterElements(nums []int) []int {
	stack := Stack{}
	stack.Init()
	result := make([]int, len(nums))

	for i, v := range nums {
		if stack.IsEmpty() {
			stack.Push(i)
			continue
		}
		for !stack.IsEmpty() && nums[stack.Top()] < v {
			result[stack.Pop()] = v
		}
		stack.Push(i)
	}
	for _, v := range nums {
		for !stack.IsEmpty() && nums[stack.Top()] < v {
			result[stack.Pop()] = v
		}
	}
	for !stack.IsEmpty() {
		result[stack.Pop()] = -1
	}

	return result
}

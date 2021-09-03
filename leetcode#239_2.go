package main

type Element struct {
	Value int
	Index int
}

type MaxStack struct {
	Values []Element
}

func (s *MaxStack) Top() Element {
	result := s.Values[0]
	return result
}

func (s *MaxStack) Pop() Element {
	result := s.Values[0]
	s.Values = s.Values[1:]
	return result
}

func (s *MaxStack) Push(e Element) {
	index := len(s.Values) - 1
	for index >= 0 && s.Values[index].Value < e.Value {
		index -= 1
	}
	if index == len(s.Values) - 1 {
		s.Values = append(s.Values, e)
	} else if index >= 0 {
		s.Values[index + 1] = e
		s.Values = s.Values[: index + 2]
	} else {
		s.Values[0] = e
		s.Values = s.Values[:1]
	}

}

func maxSlidingWindow(nums []int, k int) []int {
	stack := MaxStack{}
	result := make([]int, len(nums) - k + 1)
	i := 0
	for ; i < k; i ++ {
		stack.Push(Element{
			Value: nums[i],
			Index: i,
		})
	}
	for ; i < len(nums); i ++ {
		for {
			value := stack.Top()
			if value.Index >= i - k {
				result[i - k] = value.Value
				break
			}
			stack.Pop()
		}
		stack.Push(Element{
			Value: nums[i],
			Index: i,
		})
	}
	for {
		value := stack.Top()
		if value.Index >= i - k {
			result[i - k] = value.Value
			break
		}
		stack.Pop()
	}

	return result
}

func main() {
	maxSlidingWindow([]int{1,-1}, 1)
}
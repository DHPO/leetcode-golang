package main

type stack struct {
	Value []byte
}

func (s *stack) Push(b byte) {
	s.Value = append(s.Value, b)
}

func (s *stack) Pop() byte {
	defer func() {s.Value = s.Value[: len(s.Value) - 1]} ()
	return s.Value[len(s.Value) - 1]
}

func (s *stack) isEmpty() bool {
	return len(s.Value) == 0
}

func match(b1, b2 byte) bool {
	return (b1 == '(' && b2 == ')') || (b1 == '[' && b2 == ']') || (b1 == '{' && b2 == '}')
}

func isStart(b byte) bool {
	return b == '(' || b == '[' || b == '{'
}

func isValid(s string) bool {
	stk := stack{}
	for _, c := range []byte(s) {
		if isStart(c) {
			stk.Push(c)
		} else {
			if stk.isEmpty() || !match(stk.Pop(), c) {
				return false
			}
		}
	}
	return stk.isEmpty()
}

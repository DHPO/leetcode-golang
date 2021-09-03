package main

type ValueStack struct {
	values []int
}

func (s *ValueStack) Push(v int) {
	s.values = append(s.values, v)
}

func (s *ValueStack) Pop() int {
	result := s.values[len(s.values) - 1]
	s.values = s.values[: len(s.values) - 1]
	return result
}

func (s ValueStack) Len() int {
	return len(s.values)
}

func (s *ValueStack) PopFront() int {
	result := s.values[0]
	s.values = s.values[1:]
	return result
}

type OpStack struct {
	values []byte
}

func (s *OpStack) Push(v byte) {
	s.values = append(s.values, v)
}

func (s *OpStack) Pop() byte {
	result := s.values[len(s.values) - 1]
	s.values = s.values[: len(s.values) - 1]
	return result
}

func (s *OpStack) PopFront() byte {
	result := s.values[0]
	s.values = s.values[1:]
	return result
}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func isOp(c byte) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

func isP(c byte) bool {
	return c =='(' || c == ')'
}

func calculateFrom(s []byte, offset int) (int, int) {
	idx := offset
	readNum := func() int {
		value := 0
		for idx < len(s) {
			c := s[idx]
			if c == '(' {
				value, idx = calculateFrom(s, idx + 1)
				if s[idx] != ')' {
					panic("should be ')'")
				}
				idx += 1
				return value
			} else if isNum(c) {
				value = value * 10 + int(c - '0')
				idx += 1
			} else {
				return value
			}
		}
		return value
	}

	values := ValueStack{values: []int{}}
	ops := OpStack{values: []byte{}}

	reduce := func () int {
		value := values.PopFront()
		for values.Len() > 0 {
			op := ops.PopFront()
			value2 := values.PopFront()
			if op == '+' {
				value = value + value2
			} else {
				value = value - value2
			}
		}
		return value
	}

	for idx < len(s) {
		c := s[idx]
		if isNum(c) || c == '(' {
			values.Push(readNum())
		} else if isOp(c) {
			idx += 1
			if c == '+' || c == '-' {
				ops.Push(c)
			} else {
				prev := values.Pop()
				next := readNum()
				if c == '*' {
					values.Push(prev * next)
				} else {
					values.Push(prev / next)
				}
			}
		} else if c == ')' {
			break
		}
	}

	return reduce(), idx
}

func calculate(s string) int {
	val, _ := calculateFrom([]byte(s), 0)
	return val
}

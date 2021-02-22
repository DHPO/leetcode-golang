package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	stack := []int{}
	for ; head != nil; head = head.Next {
		stack = append(stack, head.Val)
	}
	length := len(stack)
	result := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		result[length - 1 - i] = stack[i]
	}
	return result
}

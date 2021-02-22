package main

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := 0
	node := &ListNode{}
	head := node
	for l1 != nil && l2 != nil {
		result = l1.Val + l2.Val + result / 10
		node.Next = &ListNode{Val: result % 10}
		node = node.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	var l *ListNode
	if l1 != nil {
		l = l1
	} else if l2 != nil {
		l = l2
	}
	for l != nil {
		result = l.Val + result / 10
		node.Next = &ListNode{Val: result % 10}
		node = node.Next
		l = l.Next
	}
	if result >= 10 {
		node.Next = &ListNode{Val: 1}
	}
	return head.Next
}
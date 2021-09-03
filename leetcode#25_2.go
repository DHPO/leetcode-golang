package main

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseGroup(head *ListNode, prevNode *ListNode, nextNode *ListNode) *ListNode {
	currTail := head
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	curr := head.Next

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}

	prevNode.Next = prev
	currTail.Next = nextNode

	return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	shadowHead := &ListNode{Next: head}
	start, end := head, head
	prev := shadowHead
	for start != nil {
		i := 1
		for i = 1; i < k && end.Next != nil; i ++ {
			end = end.Next
		}
		if i < k {
			break
		}
		next := end.Next
		end.Next = nil
		reverseGroup(start, prev, next)
		prev = start
		start, end = next, next
	}
	return shadowHead.Next
}
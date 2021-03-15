package main

type ListNode struct {
    Val int
    Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reorderList(head *ListNode)  {
	if head == nil {
		return
	}
	ptr1, ptr2 := head, head
	for {
		if ptr2.Next != nil {
			ptr2 = ptr2.Next.Next
		}
		if ptr2 == nil {
			return
		}
		ptr1 = ptr1.Next
	}
	lastHalf := ptr1.Next
	ptr1.Next = nil
	l2 := reverse(lastHalf)
	l1 := head
	for l1 != nil || l2 != nil {
		if l1 != nil {
			l1next := l1.Next
			l1.Next = l2
			l1 = l1next
		}
		if l2 != nil {
			l2next := l2.Next
			l2.Next = l1
			l2 = l2next
		}
	}
}
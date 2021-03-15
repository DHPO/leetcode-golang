package main

type ListNode struct {
    Val int
    Next *ListNode
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	tail := head
	var prev *ListNode = nil
	curr := tail
	for {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		if curr == nil {
			break
		}
	}
	return prev, tail
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
    if k <= 1 {
        return head
    }
	dummyHead := &ListNode{
		Next: head,
	}
	prev := dummyHead
	for {
		curr := prev.Next
		count := 1
		for curr != nil && curr.Next != nil && count < k {
			curr = curr.Next
			count ++
		}
		if count < k {
			break
		}
		tail := curr
		next := tail.Next
		tail.Next = nil
		head, tail = reverse(prev.Next)
		prev.Next = head
		tail.Next = next
		prev = tail
	}
	return dummyHead.Next
}
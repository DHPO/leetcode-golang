package main

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeList(head1, head2 *ListNode) (head, tail *ListNode) {
	head = &ListNode{}
	curr := head
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			curr.Next = head1
			curr = head1
			head1 = head1.Next
		} else {
			curr.Next = head2
			curr = head2
			head2 = head2.Next
		}
	}
	if head1 == nil {
		curr.Next = head2
	} else {
		curr.Next = head1
	}

	for curr.Next != nil {
		curr = curr.Next
	}

	head = head.Next
	tail = curr
	return
}

func sortList(head *ListNode) (*ListNode) {
	if head == nil {
		return nil
	}
	length := 0
	head = &ListNode{
		Next: head,
	}
	curr := head
	for curr.Next != nil {
		length += 1
		curr = curr.Next
	}
	for group := 1; group < 2 * length; group *= 2 {
		prev := head
		for {
			head1 := prev.Next
			tail1 := head1
			for count := 1; tail1 != nil && count < group; count ++ {
				tail1 = tail1.Next
			}
			if tail1 == nil {
				break
			}
	
			head2 := tail1.Next
			tail1.Next = nil
			tail2 := head2
	
			for count := 1; tail2 != nil && count < group; count ++ {
				tail2 = tail2.Next
			}
			var next *ListNode
			if tail2 == nil {
				next = nil
			} else {
				next = tail2.Next
				tail2.Next = nil
			}
			sortedHead, sortedTail := mergeList(head1, head2)
			prev.Next = sortedHead
			prev = sortedTail
			sortedTail.Next = next
			if next == nil {
				break
			}
		}
	}
	return head.Next
}
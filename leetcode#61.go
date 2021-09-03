package main

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || k == 0 {
        return head
    }

	fastPtr := head
	count := 0
	for fastPtr != nil {
        fastPtr = fastPtr.Next
		count += 1
		if count == k {
			break
		}
	}
	if k > count {
		k = k % count
        return rotateRight(head, k)
	}

	slowPtr := head
    if fastPtr == nil {
        return head
    }
	for fastPtr.Next != nil {
		fastPtr = fastPtr.Next
		slowPtr = slowPtr.Next
	}

	result := slowPtr.Next
	slowPtr.Next = nil
	fastPtr.Next = head

	return result
}
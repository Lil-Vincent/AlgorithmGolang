package week2

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	for p := dummy; p.Next != nil && p.Next.Next != nil; {
		a := p.Next
		b := a.Next
		p.Next = b
		a.Next = b.Next
		b.Next = a
		p = a
	}

	return dummy.Next
}

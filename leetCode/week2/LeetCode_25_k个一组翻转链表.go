package week2

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	for p := dummy; ; {
		q := p
		for i := 0; i < k && q != nil; i++ {
			q = q.Next
		}
		if q == nil {
			break
		}
		a := p.Next
		b := a.Next
		for i := 0; i < k-1; i++ {
			c := b.Next
			b.Next = a
			a = b
			b = c
		}
		c := p.Next
		p.Next = a
		c.Next = b
		p = c.Next
	}
	return dummy.Next
}

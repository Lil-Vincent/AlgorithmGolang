package week2

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	n := 0
	for p := dummy; p != nil; p = p.Next {
		n++
	}

	p := dummy
	for i := 0; i < n-k-1; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next

	return dummy.Next
}

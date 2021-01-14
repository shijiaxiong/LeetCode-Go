package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}

	p1 := dummy1
	p2 := dummy2

	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}

		head = head.Next
	}

	// 拼接
	p1.Next = dummy2.Next
	p2.Next = nil

	return dummy1.Next
}

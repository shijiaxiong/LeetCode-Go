package main

type ListNode struct {
	Val  int
	Next *ListNode
}
// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := reverseList(head.Next)

	head.Next.Next = head

	head.Next = nil

	return cur
}

// 双指针
func reverseList0(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		temp := head.Next

		head.Next = prev
		prev = head

		head = temp
	}

	return prev
}
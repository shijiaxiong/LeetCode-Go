package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head

	count := 0

	for cur != nil && count != k {
		cur = cur.Next
		count++
	}

	if count == k {
		cur = reverseKGroup(cur, k)
		for count != 0 {
			count--
			temp := head.Next
			head.Next = cur
			cur = head
			head = temp
		}

		head = cur

	}

	return head
}

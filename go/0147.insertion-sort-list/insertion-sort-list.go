package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/insertion-sort-list/solution/147-kao-cha-lian-biao-zong-he-cao-zuo-xiang-jie-by/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{}
	cur := head
	pre := dummy

	for cur != nil {
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}

		next := cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		pre = dummy
		cur = next
	}

	return dummy.Next
}

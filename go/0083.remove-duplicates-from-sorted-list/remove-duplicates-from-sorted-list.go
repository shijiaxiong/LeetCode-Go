package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 排序链表,单指针
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/solution/hua-jie-suan-fa-83-shan-chu-pai-xu-lian-biao-zhong/
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

// 双指针
// 并没有单指针简单
func deleteDuplicates0(head *ListNode) *ListNode {
	cur := head
	next := &ListNode{}

	for cur != nil && cur.Next != nil {
		next = cur.Next

		for next != nil {
			if cur.Val == next.Val {
				cur.Next = next.Next
			}
			next = next.Next
		}

		cur = cur.Next
	}

	return head
}

// 递归
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head.Next = deleteDuplicates(head.Next)

	if head.Val == head.Next.Val {
		head = head.Next
	}

	return head
}

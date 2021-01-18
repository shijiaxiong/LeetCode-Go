package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 快慢指针拆分链表
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmp := slow.Next
	slow.Next = nil
	left := sortList(head)  // 左边链表
	right := sortList(tmp) // 右边链表

	return merge(left, right)
}

func merge(left, right *ListNode) *ListNode {
	// 两个有序链表的合并
	dummy := &ListNode{}
	res := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			dummy.Next = left
			left = left.Next
		} else {
			dummy.Next = right
			right = right.Next
		}
		dummy = dummy.Next
	}
	if left != nil {
		dummy.Next = left
	}

	if right != nil {
		dummy.Next = right
	}

	return res.Next
}

// 递归
func merge2(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = merge2(list1.Next, list2)
		return list1
	}

	list2.Next = merge2(list1, list2.Next)
	return list2
}
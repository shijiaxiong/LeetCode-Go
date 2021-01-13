package main

//type ListNode utils.ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func removeNthFromEnd0(head *ListNode, n int) *ListNode {
	// 设置一个0的空节点 避免了待删除节点是head的情况
	pre := &ListNode{}
	pre.Next = head

	start, end := pre, pre

	// 先移动一个指针使得start end之间相隔n-1
	for i := 0; i < n; i++ {
		start = start.Next
	}

	// 同时移动start、end，直到start到队尾
	for start.Next != nil {
		start = start.Next
		end = end.Next
	}

	// 跳过要删除的节点
	end.Next = end.Next.Next

	return pre.Next
}

// 不设置空节点的双指针写法
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	start, end := head, head
	for i := 0; i < n; i++ {
		start  = start.Next
	}

	// 需要判断先走的指针是否走过了链表的尾部，这个时候删除的元素的head,所以返回head.next
	if start == nil {
		return head.Next
	}

	for start.Next != nil {
		start = start.Next
		end = end.Next
	}

	end.Next = end.Next.Next

	return head
}

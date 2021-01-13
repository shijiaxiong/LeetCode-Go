package main


//type ListNode utils.ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{}
	pre.Next = head

	start,end := pre,pre

	// 先移动一个指针使得start end之间相隔n
	for i:=0;i<n;i++{
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

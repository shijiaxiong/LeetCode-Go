package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 类似第19题
// 双指针
func rotateRight(head *ListNode, k int) *ListNode {
	// head为nil、head只有一个元素、k为0直接返回
	if head == nil || k == 0 || head.Next == nil {
		return head
	}

	// 计算链表的长度
	temp := head
	count := 0
	for temp != nil {
		temp = temp.Next
		count++
	}

	// k 大于链表长度取模
	k = k % count
	if k == 0 {
		return head
	}

	// 双指针
	start, end := head, head

	for i := 0; i < k; i++ {
		start = start.Next
	}

	for start.Next != nil {
		start = start.Next
		end = end.Next
	}

	res := end.Next
	// 不置为空就会形成环
	end.Next = nil
	start.Next = head

	return res
}

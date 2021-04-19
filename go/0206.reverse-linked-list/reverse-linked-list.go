package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func reverseList(head *ListNode) *ListNode {
	// 递归终止条件是当前为空，或者下一个节点为空
	if head == nil || head.Next == nil {
		// 返回的head是cur的值
		return head
	}

	// 这里的cur就是链表的最后一个节点
	cur := reverseList(head.Next)

	//如果链表是 1->2->3->4->5，那么此时的cur就是5
	//而head是4，head的下一个是5，下下一个是空
	//所以head.next.next 就是5->4
	head.Next.Next = head

	head.Next = nil

	return cur
}

// 双指针
// 使用了head代替cur，不熟悉不好理解 不推荐
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

// prev cur 同时移动
// https://leetcode-cn.com/problems/reverse-linked-list/solution/fan-zhuan-lian-biao-shuang-zhi-zhen-di-gui-yao-mo-/
func reverseList1(head * ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
}
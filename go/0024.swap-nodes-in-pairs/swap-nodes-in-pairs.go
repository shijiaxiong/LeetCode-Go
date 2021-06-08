package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/swap-nodes-in-pairs/solution/dong-hua-yan-shi-24-liang-liang-jiao-huan-lian-bia/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//假设链表是 1->2->3->4
	//这句就先保存节点2
	next := head.Next

	//继续递归，处理节点3->4
	//当递归结束返回后，就变成了4->3
	//于是head节点就指向了4，变成1->4->3
	head.Next = swapPairs(next.Next)

	//将2节点指向1
	next.Next = head

	return next
}

// 迭代
func swapPairs1(head *ListNode) *ListNode {
	prev := &ListNode{}
	prev.Next = head

	temp := prev

	for temp.Next != nil && temp.Next.Next != nil {
		start := temp.Next
		end := temp.Next.Next

		temp.Next = end

		start.Next = end.Next
		end.Next = start
		temp = start
	}

	return prev.Next
}
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归解法：https://leetcode-cn.com/problems/reverse-linked-list-ii/solution/bu-bu-chai-jie-ru-he-di-gui-di-fan-zhuan-lian-biao/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}

	// 前进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

var successor *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 记录第 n + 1 个节点
		successor = head.Next
		return head
	}
	// 以 head.next 为起点，需要反转前 n - 1 个节点
	last := reverseN(head.Next, n-1)

	head.Next.Next = head
	// 让反转之后的 head 节点和后面的节点连起来
	head.Next = successor
	return last

}

// 双指针
// https://leetcode-cn.com/problems/reverse-linked-list-ii/solution/java-shuang-zhi-zhen-tou-cha-fa-by-mu-yi-cheng-zho/
func reverseBetween0(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	g := dummy
	p := dummy.Next

	// 移动到m点
	for i := 0; i < m-1; i++ {
		g = g.Next
		p = p.Next
	}

	for i := 0; i < n-m; i++ {
		remove := p.Next
		p.Next = p.Next.Next
		remove.Next = g.Next
		g.Next = remove
	}

	return dummy.Next
}

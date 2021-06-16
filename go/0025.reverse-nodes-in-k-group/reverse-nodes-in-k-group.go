package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/kge-yi-zu-fan-zhuan-lian-biao-by-powcai/
// 递归
func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	count := 0

	//  计数
	for cur != nil && count != k {
		cur = cur.Next
		count++
	}

	if count == k {
		//  下一组k个节点进行反转
		cur = reverseKGroup(cur, k)

		for count != 0 {
			count--
			temp := head.Next
			head.Next = cur
			cur = head
			head = temp
		}

		//  执行完循环，cur 是当前k个节点的头
		head = cur
	}

		return head
}
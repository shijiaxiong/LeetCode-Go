package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 题解：https://www.cxyxiaowu.com/6857.html
// 堆相关：https://ieevee.com/tech/2018/01/29/go-heap.html
//由于 heap 的大小为始终为 k ，而每次插入的复杂度是 logk ，一共插入了 nk 个节点。时间复杂度为 O(nklogk)，空间复杂度为O(k)
func mergeKLists(lists []*ListNode) *ListNode {
	h := HeapMin{}
	heap.Init(&h)

	node := &ListNode{}
	cur := node
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(&h, lists[i])
		}
	}

	for h.Len() > 0 {
		tmp := heap.Pop(&h).(*ListNode)
		cur.Next = tmp
		cur = cur.Next
		if node.Next != nil {
			heap.Push(&h, tmp.Next)
		}
	}

	return node.Next
}

type HeapMin []*ListNode

func (h HeapMin) Len() int {
	return len(h)
}

func (h HeapMin) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *HeapMin) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *HeapMin) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

func (h *HeapMin) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

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


// 归并排序
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return mergeSort(lists, 0, len(lists) - 1)
}

func mergeSort(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	mid := left + (right - left) / 2
	l1 := mergeSort(lists, left, mid)
	l2 := mergeSort(lists, mid + 1, right)

	return mergeList(l1, l2)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeList(l1, l2.Next)
		return l2
	}
}
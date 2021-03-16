package main

import (
	"container/heap"
)

// 快排序
// 利用快排的性质，基准值左边比它小，右边比它大
func getLeastNumbers0(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	start, end := 0, len(arr)-1
	// 找基准值
	index := partition(arr, start, end)

	for index != k-1 {
		// 基准值比 k 大,基准值前面的数组都比K小
		if index > k-1 {
			end = index - 1
		} else {
			// 基准比K小
			start = index + 1
		}

		index = partition(arr, start, end)
	}

	return arr[:k]
}

func partition(arr []int, l, h int) int {
	mid := arr[l]
	for l < h {
		for l < h && arr[h] >= mid {
			h--
		}
		arr[l] = arr[h]

		for l < h && arr[l] <= mid {
			l++
		}
		arr[h] = arr[l]
	}

	arr[l] = mid
	return l
}

// 堆排序
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	h := &IHeap{}
	heap.Init(h)

	for _, v := range arr {
		if h.Len() < k {
			heap.Push(h, v)
		} else {
			if (*h)[0] > v {
				heap.Pop(h)
				heap.Push(h, v)
			}
		}
	}

	var res []int
	for h.Len()>0{
		res = append(res, heap.Pop(h).(int))
	}
	return res

}

type IHeap []int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IHeap) Pop(v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

package main

import (
	"container/heap"
)

// 快排序
// 利用快排的性质，基准值左边比它小，右边比它大
// 时间复杂度：O(n)
func getLeastNumbers0(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	left, right := 0, len(arr)-1
	// 找基准值
	index := partition(arr, left, right)

	for index != k-1 {
		// 基准值比 k 大,基准值前面的数组都比K小
		if index > k-1 {
			right = index - 1
		} else {
			// 基准比K小
			left = index + 1
		}

		index = partition(arr, left, right)
	}

	return arr[:k]
}

func partition(arr []int, left, right int) int {
	pivot := arr[left]

	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left] = arr[right]

		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}

	arr[left] = pivot
	return left
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

// 手撕堆排
func getLeastNumbers2(arr []int, k int) []int {
	return heapSort(arr,k)
}

func heapSort(nums []int, k int) []int {
	lens := len(nums)
	//  建堆
	for i := lens/2 - 1;i >=0;i-- {
		heapify(nums, lens, i)
	}

	//  排序
	for i := lens - 1; i >0; i-- {
		nums[0], nums[i] = nums[i],nums[0]

		heapify(nums, i, 0)
	}

	return nums[:k]
}

func heapify(nums []int, n, i int) {
	largest := i
	lson := i * 2 + 1
	rson := i * 2 + 2

	if lson < n && nums[largest] < nums[lson] {
		largest = lson
	}

	if rson < n && nums[largest] < nums[rson] {
		largest = rson
	}

	if largest != i {
		nums[i], nums[largest]  = nums[largest], nums[i]
		heapify(nums, n, largest)
	}
}
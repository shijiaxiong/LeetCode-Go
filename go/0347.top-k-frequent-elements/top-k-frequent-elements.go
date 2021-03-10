package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)

	// 统计每个数字出现的次数
	rec := make(map[int]int, len(nums))
	for _, n := range nums {
		rec[n]++
	}

	//	对出现次数进行排序
	counts := make([]int, 0, len(rec))
	for _, c := range rec {
		counts = append(counts, c)
	}

	sort.Ints(counts)

	// min 是 前 k 个高频数字的底线
	min := counts[len(counts)-k]

	for n, c := range rec {
		if c >= min {
			res = append(res, n)
		}
	}

	return res
}

// 最小堆
// k 个元素
// 时间复杂度：O(nlogk)
// 空间复杂度：O(n)
func topKFrequent0(nums []int, k int) []int {
	// 使用字段统计每个元素出现的次数
	hashMap := map[int]int{}
	for _, num := range nums {
		hashMap[num]++
	}

	h := &IHeap{}
	heap.Init(h)

	// 遍历map 用最小堆保存频率最大的k个元素
	for key, value := range hashMap {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 最小值在堆顶 所以先弹出的是最小值。也可以逐个弹出
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}

	return res
}

type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

func main() {
	topKFrequent1([]int{1, 1, 1, 2, 2, 3}, 2)
}

//
func topKFrequent1(nums []int, k int) []int {
	hashMap := map[int]int{}
	for _, num := range nums {
		hashMap[num]++
	}
	fmt.Println(hashMap)

	// value最小为1，即至少出现一次
	// bucket下标大 意味着出现频率高
	bucket := make([][]int, len(nums))
	for key, value := range hashMap {
		bucket[value - 1] = append(bucket[value - 1], key)
	}

	fmt.Println(bucket)

	var res []int
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) == 0 {
			continue
		}

		res = append(res, bucket[i]...)
		if len(res) >= k {
			break
		}
	}

	return res
}

package main

import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	//  使用字典，统计每个元素出现的次数，元素为键，元素出现的次数为值
	hash := map[int]int{}
	for i := range nums {
		hash[nums[i]]++
	}

	//桶排序
	//将频率作为数组下标，对于出现频率不同的数字集合，存入对应的数组下标
	//  出现的次数是key,key越大说明出现的频率越高
	bucket := make([][]int, len(nums))
	for key,value := range hash {
		bucket[value - 1] = append(bucket[value - 1], key)
	}

	// 倒序遍历数组获取出现顺序从大到小的排列
	res := []int{}
	for i := len(bucket) - 1; i >=0; i-- {
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

}


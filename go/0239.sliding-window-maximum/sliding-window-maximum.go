package main

import (
	"math"
)

// 双端队列 队列内的数据保持递减
func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil {
		return []int{}
	}

	// 队列中存储的是nums元素的下标
	// res 存结果
	queue, res := []int{}, []int{}

	for i, x := range nums {
		// 当队列中最左侧的元素下标 超出窗口范围的时候剔除
		if i >= k && queue[0] <= i-k {
			queue = queue[1:]
		}

		// 比当前元素值小就从栈顶弹出, 始终维护一个从大到小的队列
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= x {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)

		// 从第k - 1开始，只要i >= k - 1 可以将最左元素添加到结果
		if i >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}

	return res
}

func main() {
	//[1,3,-1,-3,5,3,6,7]
	//3
	maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}

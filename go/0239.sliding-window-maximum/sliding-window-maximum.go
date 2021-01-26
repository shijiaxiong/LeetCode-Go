package main

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil {
		return []int{}
	}

	// window存下标
	// res 存结果
	window, res := []int{}, []int{}

	for i, x := range nums {
		// 当window装满后再开始往后挪动窗口
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}

		// 逐个删除最大值左边的各个元素
		for len(window) > 0 && nums[window[len(window)-1]] <= x {
			window = window[:len(window)-1]
		}

		window = append(window, i)

		// 从第k - 1开始，只要i >= k - 1 可以将最左元素添加到结果
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}

	return res
}

func main() {
	//[1,3,-1,-3,5,3,6,7]
	//3
	maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}

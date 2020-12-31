package problem123

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(myrob(nums, 0, n-2), myrob(nums, 1, n-1))
}

func myrob(nums []int, start, end int) int {
	pre, cur := nums[start], max(nums[start], nums[start+1])

	for i := start + 2; i <= end; i++ {
		temp := cur
		cur = max(cur, nums[i]+pre)
		pre = temp
	}

	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

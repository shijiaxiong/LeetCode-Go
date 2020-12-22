package problem0198

func rob(nums []int) int {
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			a = max(a+nums[i],b)
		} else {
			b = max(b + nums[i], a)
		}
	}

	return max(a,b)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

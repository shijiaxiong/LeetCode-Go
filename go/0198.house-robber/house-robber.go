package problem0198

// 动态规划
// 动态规划方程：dp[i] = max(dp[i-1], dp[i-2] + nums[i])
func rob1(nums []int)int{
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// dp[i] 代表抢 nums[0...i] 房子的最大价值
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}

// 动态规划优化辅助空间
func rob2(nums []int)int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	pre , cur := 0,0

	for _,v := range nums{
		temp := cur
		cur = max(cur,v+pre)
		pre = temp
	 }

	 return cur
}

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

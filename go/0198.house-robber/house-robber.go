package problem0198

// 动态规划
// https://leetcode-cn.com/problems/house-robber/solution/dong-tai-gui-hua-jie-ti-si-bu-zou-xiang-jie-cjavap/
// 动态规划方程：dp[i] = max(dp[i-1], dp[i-2] + nums[i])
func rob1(nums []int)int{
	if len(nums) == 0 {
		return 0
	}
	// dp[i] 表示盗取第i间房的价值
	dp := make([]int, len(nums) + 1)
	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i <= len(nums); i++ {
		// 第i间房子的价值是nums[i-1]
		dp[i] = max(dp[i-1], dp[i-2] + nums[i - 1])
	}

	return dp[len(nums)]
}

// 动态规划优化辅助空间
func rob2(nums []int)int {
	if len(nums) == 0 {
		return 0
	}

	prev := 0
	curr := nums[0]

	for i := 2; i <= len(nums); i++ {
		temp := max(curr, prev + nums[i - 1])
		prev = curr
		curr = temp
	}

	return curr
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

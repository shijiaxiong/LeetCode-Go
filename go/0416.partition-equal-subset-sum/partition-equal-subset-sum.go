package main

// 动态规划  转换成0-1背包问题
// 给一个可以装sum(nums)/2的背包、len(nums)件物品，物品的重量为nums[i] 能否正好把背包装满
func canPartition(nums []int) bool {
	length := len(nums)

	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make([][]bool, length+1)

	for i := 0; i <= length; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true // 单独的一个数
	}

	for i := 1; i <= length; i++ {
		for j := 1; j <= target; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[length][target]
}

// 状态压缩
func canPartition0(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make([]bool, target+1)

	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := target; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}

	return dp[target]
}

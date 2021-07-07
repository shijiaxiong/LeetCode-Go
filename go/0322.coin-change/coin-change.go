package problem0322

// 动态规划
//https://leetcode-cn.com/problems/coin-change/solution/dong-tai-gui-hua-shi-yong-wan-quan-bei-bao-wen-ti-/
// 背包问题
// 动态规划
// dp[i] = min(dp[i], dp[i-coin[i]] + 1)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 理解 dp[0] = 0 的合理性，单独一枚硬币如果能够凑出面值，符合最优子结构
	// dp[0] - 要凑的面值，0 - 硬币数
	// 凑够dp[i]数量的金钱需要的硬币数
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		//注意：因为要比较的是最小值，这个不可能的值就得赋值成为一个最大值
		dp[i] = amount + 1
		for _, coin := range coins {
			// 面值不超过要够凑的金额 amount
			// 当前要凑够i的硬币数(之前已经赋予不可能的最大值) 就是dp[i-c] + 1
			if coin <= i && dp[i] > dp[i-coin]+1 {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	// 当前金额不能被给定硬币表示
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

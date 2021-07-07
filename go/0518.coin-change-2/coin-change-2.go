package main

//[1,3,2,4,7,10]
//

// 动态规划
// 状态转移方程
// if 金额数大于硬币
//    DP[k][i] = DP[k-1][i] + DP[k][i-k]
//else
//    DP[k][i] = DP[k-1][i]
// 背包问题：背包最大容量amount,有len(coins)类物品，每个物品的价值coins[i], 每类物品数量无限
func change(amount int, coins []int) int {
	// 组成的矩形 垂直方向是硬币数；水平方向是金额
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	// 初始化数组
	for i := 0; i < len(coins)+1; i++ {
		for j := 0; j < amount+1; j++ {
			dp[i][j] = 0
		}
	}

	// 初始化基本状态
	for i := 0; i < len(coins)+1; i++ {
		dp[i][0] = 1
	}

	for k := 1; k <= len(coins); k++ {
		for i := 1; i <= amount; i++ {
			if i >= coins[k-1] {
				dp[k][i] = dp[k][i-coins[k-1]] + dp[k-1][i]
			} else {
				dp[k][i] = dp[k-1][i]
			}
		}
	}
	return dp[len(coins)][amount]
}

// 动态规划空间优化
// dp[i] 代表装满容量为i的背包有几种硬币组合
// dp[i] = dp[i]+ dp[i-k];dp[i] 表示对于第k个硬币能凑的组合数
func change0(amount int, coins []int) int {
	// 对特殊情况处理可以减少耗时
	if amount == 0 {
		return 1
	}

	if len(coins) == 0 {
		return 0
	}

	//dp := make([]int, 5001) // 根据题目要求 直接给定数组大小可以更省空间
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, c := range coins {
		for i := 1; i <= amount; i++ {
			if c <= i { // 硬币面额不可以大于金额
				dp[i] = dp[i] + dp[i-c]
			}
		}
	}

	return dp[amount]
}

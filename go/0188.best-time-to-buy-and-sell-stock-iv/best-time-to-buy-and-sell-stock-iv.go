package main

// 动态规划
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	length := len(prices)

	// k 超过length/2(可盈利的最大交易次数)
	if k >= length/2 {
		return maxProfitA(prices)
	}

	dp := make([][][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 1; i <= k; i++ {
		dp[0][i][0] = 0
		dp[0][i][1] = -prices[0]
	}

	for i := 1; i < length; i++ {
		for j := k; j > 0; j-- {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[length-1][k][0]
}

func maxProfitA(prices []int) int {
	length := len(prices)

	if length == 0 {
		return 0
	}

	dp := make([][2]int, length)

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[length-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 动态规划 优化
func maxProfit0(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	length := len(prices)

	// k 超过length/2(可盈利的最大交易次数)
	if k >= length/2 {
		return maxProfitB(prices)
	}

	dp := make([][2]int, k+1)

	for i := 1; i <= k; i++ {
		dp[i][0] = 0
		dp[i][1] = -prices[0]
	}

	for i := 1; i < length; i++ {
		for j := k; j > 0; j-- {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
		}
	}

	return dp[k][0]
}

func maxProfitB(prices []int) int {
	length := len(prices)

	if length == 0 {
		return 0
	}

	profit0 := 0
	profit1 := -prices[0]
	for i := 1; i < length; i++ {
		profit0 = max(profit0, profit1+prices[i])
		profit1 = max(profit1, profit0-prices[i])
	}

	return profit0
}

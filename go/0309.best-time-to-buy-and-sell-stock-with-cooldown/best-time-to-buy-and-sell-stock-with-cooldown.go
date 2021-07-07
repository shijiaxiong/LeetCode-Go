package main

// 动态规划
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/solution/dong-tai-gui-hua-by-liweiwei1419-5/
func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}

	dp := make([][3]int, length)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	// dp[i][0]: 手上不持有股票，并且今天不是由于卖出股票而不持股，我们拥有的现金数
	// dp[i][1]: 手上持有股票时，我们拥有的现金数
	// dp[i][2]: 手上不持有股票，并且今天是由于卖出股票而不持股，我们拥有的现金数
	for i := 1;  i < len(prices); i++ {
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][2])
		dp[i][1] = max(dp[i - 1][1], dp[i-1][0] - prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
	}

	return max(dp[length - 1][0], dp[length - 1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit0(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}

	prevProfit0 := 0
	profit0 := 0
	profit1 := -prices[0]

	for i := 1; i < length; i++ {
		nextProfit0 := max(profit0, profit1+prices[i])
		nextProfit1 := max(profit1, prevProfit0-prices[i])
		prevProfit0 = profit0
		profit0 = nextProfit0
		profit1 = nextProfit1

	}

	return profit0
}

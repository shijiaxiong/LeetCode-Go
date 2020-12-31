package main

// 动态规划
func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}

	dp := make([][2]int, length)

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		if i >= 2 {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], 0-prices[i])
		}
	}

	return dp[length-1][0]
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

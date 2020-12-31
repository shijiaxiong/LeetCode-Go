package problem0121

import (
	"math"
)

// 贪心算法
func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if maxProfit < prices[i]-minPrice {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

// 动态规划
// 证明 结束时持有 0 份股票的收益一定大于持有 1 份股票的收益？
// 如果结束时持有 1 份股票，可能有两种情况，一是在最后一天买入，二是在之前买入。
// 如果是第一种情况，则最后一天选择不买入，收益一定更高，因为不用付出购买股票的钱。
// 如果是第二种情况，则一定可以在最后一天卖出，卖出操作一定会增加收益。
// 因此，无论是哪种情况，持有 0 份股票的收益一定大于持有 1 份股票的收益。
func maxProfit0(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	length := len(prices)

	dp := make([][2]int, length)

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[length-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 动态规划 优化存储空间
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	length := len(prices)

	profit0 := 0
	profit1 := -prices[0]

	for i := 1; i < length; i++ {
		profit0 = max(profit0, profit1+prices[i])
		profit1 = max(profit1, -prices[i])
	}

	return profit0
}
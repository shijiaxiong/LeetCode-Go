package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	//  dp[i][j][k] j:交易次数，k 持有的股票
	//  第 2 维的 0 没有意义，1 表示交易进行了 1 次，2 表示交易进行了 2 次 ， 为了使得第 2 维的数值 1 和 2 有意义，这里将第 2 维的长度设置为 3
	dp := make([][3][2]int, len(prices))

	// dp[0][0][0] = 0
	// dp[0][0][1] = 0 不可能状态
	// dp[0][1][0] // 一次交易不持股，不可能
	dp[0][1][1] = -prices[0]
	// dp[0][2][0] // 两次交易不持股，不可能
	dp[0][2][1] = math.MinInt64 // 两次交易依然持股，也可以给个-prices[0]

	for i := 1; i < len(prices); i++ {
		//  先持股 再卖出
		dp[i][1][1] = max(dp[i-1][1][1], -prices[i])
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
		dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - prices[i])
		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + prices[i])
	}

	return max(dp[len(prices) - 1][1][0], dp[len(prices) - 1][2][0])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func maxProfit0(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	length := len(prices)

	profitOne0 := 0
	profitOne1 := -prices[0]
	profitTwo0 := 0
	profitTwo1 := -prices[0]

	for i := 1; i < length; i++ {
		profitTwo0 = max(profitTwo0, profitTwo1+prices[i])
		profitTwo1 = max(profitTwo1, profitOne0-prices[i])
		profitOne0 = max(profitOne0, profitOne1+prices[i])
		profitOne1 = max(profitOne1, -prices[i])
	}

	return profitTwo0
}

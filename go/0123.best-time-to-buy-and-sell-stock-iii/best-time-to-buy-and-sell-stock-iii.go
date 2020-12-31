package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	length := len(prices)

	dp := make([][3][2]int, length)

	dp[0][1][0] = 0
	dp[0][1][1] = -prices[0]
	dp[0][2][0] = 0
	dp[0][2][1] = -prices[0]

	for i := 1; i < length; i++ {
		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
		dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-prices[i])
	}

	return dp[length-1][2][0]
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

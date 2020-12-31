package problem0122

//贪心算法
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	res := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

// 贪心算法
func maxProfit0(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}

	ans := 0
	inHand := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < inHand {
			inHand = prices[i]
		} else if prices[i] > inHand {
			ans += prices[i] - inHand
			inHand = prices[i]
		}
	}

	return ans
}

// 动态规划
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	length := len(prices)

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

// 动态规划降维
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	length := len(prices)
	profit0 := 0
	profit1 := -prices[0]
	for i := 1; i < length; i++ {
		profit0 = max(profit0, profit1+prices[i])
		profit1 = max(profit1, profit0-prices[i])
	}

	return profit0
}

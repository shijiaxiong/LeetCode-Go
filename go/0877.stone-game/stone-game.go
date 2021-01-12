package main

// 动态规划
// 类似最长回文子序列思路
//dp[i][j] = max { pile[i] - dp[i+1][j] , pile[j] - dp[i][j-1] }
func stoneGame(piles []int) bool {
	length := len(piles)

	dp := make([][]int, length)
	for k := range dp {
		dp[k] = make([]int, length)
	}

	for i := 0; i < length; i++ {
		dp[i][i] = piles[i]
	}

	// 从下到上 i不包含对角线
	for i := length - 1; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			// 选i堆石子 - 在dp[i+1][j]中的得分
			// 选j堆石子 - 在dp[i+1][j-1]中的得分
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}

	// 或
	//for j:=1;j<length;j++{
	//	for i := j-1;i>=0;i--{
	//		dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
	//	}
	//}

	return dp[0][length-1] > 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

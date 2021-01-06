package main

// 动态规划
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}

	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

// 动态规划空间优化
func uniquePathsWithObstacles0(obstacleGrid [][]int) int {
	n := len(obstacleGrid[0])

	dp := make([]int, n)
	dp[0] = 1

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < n; j++ {

			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			}

			if j > 0 && obstacleGrid[i][j] == 0 {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}

	return dp[n-1]
}

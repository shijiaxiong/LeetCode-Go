package main

func minPathSum(grid [][]int) int {
	m := len(grid)    // 从上往下
	n := len(grid[0]) // 从左往右

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i][j-1] , dp[i-1][j])
		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// https://leetcode-cn.com/problems/minimum-path-sum/solution/hui-su-dao-dong-tai-gui-hua-zai-dao-kong-swk9/
func minPathSum0(grid [][]int) int {
	//  自底下上
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = grid[m-1][n-1]

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			//  最后一行(不包含最后一格)
			if i == m - 1 && j != n - 1 {
				dp[i][j] = grid[i][j] + dp[i][j+1]
			} else if i != m-1 && j == n-1 { // 最后一列 不包含最后一格
				dp[i][j] = grid[i][j] + dp[i+1][j]
			} else if i != m-1 && j != n-1 {
				dp[i][j] = grid[i][j] + min(dp[i+1][j], dp[i][j + 1])
			}
		}
	}

	return dp[0][0]
}

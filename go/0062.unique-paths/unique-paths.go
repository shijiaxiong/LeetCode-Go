package main

// 动态规划
//时间复杂度：O(m*n)
//空间复杂度：O(m * n)
func uniquePaths(m int, n int) int {
	// 按题目中m n访问不超过100
	dp := [100][100]int{}

	//dp := make([][]int, m)
	//for k := range dp {
	//	dp[k] = make([]int, n)
	//}

	// 初始化最左侧和最上层，只有一种走法
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// 动态规划空间优化 空间复杂度 O(2n)
// 因为dp[i][j] 的结果只与它上边和左边的数有关， 所以可以使用两个数组 表示上一行和下一行。
func uniquePaths0(m int, n int) int {
	pre := make([]int, n)
	cur := make([]int, n)

	for i := 0; i < n; i++ {
		pre[i] = 1
		cur[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] = cur[j-1] + pre[j]
			pre = cur
		}
	}

	//由于有pre = cur的赋值 此处也可以返回cur[n-1]
	return pre[n-1]
}

// 动态规划空间优化 O(n)
func uniquePaths1(m int, n int) int {
	cur := make([]int, n)

	for i := 0; i < n; i++ {
		cur[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] = cur[j] + cur[j-1]
		}
	}

	return cur[n-1]
}

package main

// 动态规划
// if s[i] == s[j] {
// 		dp[i][j] = dp[i+1][j-1] + 2
// } else {
// 		dp[i][j] = max(dp[i][j-1], dp[i+1][j])
// }
//https://leetcode-cn.com/problems/longest-palindromic-subsequence/solution/zi-xu-lie-wen-ti-tong-yong-si-lu-zui-chang-hui-wen/
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	// 从下往上
	for i := n - 1; i >= 0; i-- {
		// 从左往右
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

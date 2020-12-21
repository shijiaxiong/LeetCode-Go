package problem1143

// 二维矩阵的动态规划
// 画dp-table可以方便理解
func longestCommonSubsequence(text1 string, text2 string) int {

	dp := make([][]int, len(text1) + 1)

	for k, _ := range dp {
		dp[k] = make([]int, len(text2) +1)
	}

	// 含有空字符串 所有len+1
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划 更优化解答
func longestCommonSubsequence1(s1, s2 string) int {
	m, n := len(s1), len(s2)
	A, B := []byte(s1), []byte(s2)

	cur := make([]int, m+1)
	prev := make([]int, m+1)

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			rec := max(prev[i+1], cur[i])
			if A[i] == B[j] {
				cur[i+1] = max(rec, prev[i]+1)
			} else {
				cur[i+1] = rec
			}
		}
		cur, prev = prev, cur
	}
	return prev[m]
}

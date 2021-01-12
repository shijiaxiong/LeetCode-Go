package main

// 动态规划
// 找状态 一个状态就是数组的一个维度
func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	if sLen == 0 || pLen == 0 {
		return false
	}

	dp := make([][]bool, sLen+1)
	for i := range dp {
		dp[i] = make([]bool, pLen+1)
	}

	//	 base case
	dp[0][0] = true
	for j := 1; j < pLen+1; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i < sLen+1; i++ {
		for j := 1; j < pLen+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[sLen][pLen]
}

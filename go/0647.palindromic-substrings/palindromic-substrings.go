package main

// https://leetcode-cn.com/problems/palindromic-substrings/solution/647-hui-wen-zi-chuan-dong-tai-gui-hua-fang-shi-qiu/
// 动态规划
func countSubstrings(s string) int {
	length := len(s)
	count := length
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}

	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				if j - i == 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}

			if dp[i][j] {
				count++
			}
		}
	}

	return count
}

package main

func longestPalindrome(s string) string {
	length := len(s)

	if length < 2 {
		return s
	}

	maxLen, begin := 1, 0
	dp := make([][]bool, length)

	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)

		dp[i][i] = true
	}

	for j := 1; j < length; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			//	只要 dp[i][j] == true 成立，就表示子串 s[i..j] 是回文，此时记录回文长度和起始位置
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+maxLen]
}

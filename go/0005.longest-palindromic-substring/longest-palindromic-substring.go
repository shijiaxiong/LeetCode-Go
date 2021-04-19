package main

// https://leetcode-cn.com/problems/longest-palindromic-substring/solution/5-zui-chang-hui-wen-zi-chuan-dong-tai-gui-hua-jie-/
func longestPalindrome(s string) string {
	// 过滤单个字符的场景
	length := len(s)

	if length < 2 {
		return s
	}

	result := [2]int{}

	dp := make([][]bool, length)

	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)

		// 在dp数组图中符合条件的只是对称线上的点
		dp[i][i] = true
	}

	// 从右下角开始遍历，右下角往上一格开始
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				// i j相邻
				if j-i == 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}

			//	只要 dp[i][j] == true 成立，就表示子串 s[i..j] 是回文，此时记录回文长度和起始位置
			if dp[i][j] && result[1]-result[0] <= j-i {
				result[0] = i
				result[1] = j
			}
		}
	}

	return s[result[0] : result[1]+1]
}

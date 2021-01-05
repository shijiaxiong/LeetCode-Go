package main

import (
	"fmt"
)

// 动态规划
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)

	dp := make([][]int, len1+1)
	for k := range dp {
		dp[k] = make([]int, len2+1)
	}

	// 初始化：当 word2 长度为 0 时，将 word1 的全部删除即可
	for i := 1; i <= len1; i++ {
		dp[i][0] = i
	}

	// 当 word1 长度为 0 时，插入所有 word2 的字符即可
	for i := 1; i <= len2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			insert := dp[i][j-1] + 1
			replace := dp[i-1][j-1] + 1
			delete := dp[i-1][j] + 1

			dp[i][j] = min(min(insert, replace), delete)
		}
	}

	return dp[len1][len2]

}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

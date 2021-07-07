package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	imax := 0
	var stack []int

	stack = append(stack, -1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				imax = max(imax, i-stack[len(stack)-1])
			}
		}
	}

	return imax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划
// https://leetcode-cn.com/problems/longest-valid-parentheses/solution/shou-hua-tu-jie-zhan-de-xiang-xi-si-lu-by-hyj8/
func longestValidParentheses0(s string) int {
	res := 0
	//  以s[i]结尾的字符的有效长度
	dp := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i - 1] == '(' {
				if i - 2 >= 0 {
					dp[i] = dp[i-2] + 2 // dp[i-2] 存在
				} else {
					dp[i] = 2
				}
			} else if i - dp[i-1] - 1 >= 0 && s[i - dp[i-1] - 1] == '(' {   // 跨过 以s[i-1]为结尾的最长有效长度 dp[i-1]
				if i - dp[i-1] - 2 >= 0 {
					dp[i] = dp[i - dp[i-1] - 2] + dp[i-1 ] + 2
				} else {
					dp[i] = dp[i - 1] + 2
				}
			}
		}

		res = max(res, dp[i])
	}

	return res
}

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

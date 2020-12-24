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
			stack = stack[: len(stack)-1]
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

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

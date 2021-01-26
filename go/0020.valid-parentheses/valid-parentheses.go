package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []string{}

	for _, v := range s {
		cur := string(v)
		// 左括号等待匹配
		if cur == "(" || cur == "[" || cur == "{" {
			stack = append(stack, cur)
		} else {
			if len(stack) == 0 {
				return false
			}

			// 查看stack中的最后一个字符 和 当前的轮询到的字符串是否匹配
			top := stack[len(stack)-1]
			if top == "(" && cur == ")" || top == "[" && cur == "]" || top == "{" && cur == "}" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

// 执行时间较长
func isValid0(s string) bool {
	m := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	lps := []string{}

	for i := 0; i < len(s); i++ {
		p := string(s[i])

		if _, ok := m[p]; ok {

			// 都存了左括号
			lps = append(lps, p)
		} else {
			// lps 不为空 并且lps中最后一个元素是p对应的括号
			if len(lps) > 0 && m[lps[len(lps)-1]] == p {

				lps = lps[:len(lps)-1] // 剔除匹配到的字符串
				fmt.Println(lps)
			} else {
				return false
			}
		}
	}

	if len(lps) == 0 {
		return true
	}

	return false
}

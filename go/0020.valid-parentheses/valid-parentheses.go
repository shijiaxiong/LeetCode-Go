package main

import (
	"fmt"
)

func isValid(s string) bool {
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	lps := make([]byte, len(s)/2)

	fmt.Println(lps)

	for i := 0; i < len(s); i++ {
		p := s[i]

		if _,ok := m[p]; ok{
			lps = append(lps, p)
		} else {
			// lps 不为空 并且lps中最后一个元素是p对应的括号
			if len(lps) > 0 && m[lps[len(lps)-1]] == p {
				lps = lps[:len(lps)-1]	// 剔除匹配到的字符串
			} else {
				return false
			}
		}
	}

	if len(lps) == 0{
		return true
	}

	return false
}

func main() {
	fmt.Println(isValid("()"))
}

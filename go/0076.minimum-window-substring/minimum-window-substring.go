package main

import "math"

// 滑动窗口
func minWindow(s string, t string) string {
	need := [128]int{}

	for i := range t {
		need[t[i]] ++
	}

	needCnt := len(t)

	i := 0
	size := math.MaxInt32
	count := len(t)
	start := 0

	for j := 0; j < len(s); j++ {
		if need[s[i]] > 0 {
			needCnt--
		}

		need[s[i]]--

		// 窗口中已经包含t中的所有字符
		if needCnt == 0 {

			for i < j && need[s[i]] < 0 {
				need[s[i]]++
				i++
			}

			if j-i+1 < size {
				size = j - i + 1
				start = i
			}
			need[s[i]]++
			count++
		}
	}

	if size == math.MaxInt32 {
		return ""
	}

	return s[start : start+size]
}

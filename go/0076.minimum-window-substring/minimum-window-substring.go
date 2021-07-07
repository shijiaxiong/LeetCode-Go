package main

import (
	"fmt"
)

// 滑动窗口
// 参考：https://leetcode-cn.com/problems/minimum-window-substring/solution/hua-dong-chuang-kou-ji-bai-liao-100de-javayong-hu-/
func minWindow(s string, t string) string {
	//  窗口中s中各个字符出现的次数
	have := [128]int{}
	//  t中各个字符出现的次数，作为参照
	need := [128]int{}

	// 统计t中各个字符出现的次数
	for i := range t {
		need[t[i]]++
	}

	//  i, j窗口的左右指针；count 统计已有t中字母的数量
	i := 0
	j := 0
	count := 0

	// min用来统计最小字符串 赋给一个不可能的大值
	min := len(s) + 1
	res := ""

	for j < len(s) {
		have[s[j]]++

		//  s中出现了t中的字符
		if need[s[j]] > 0 && need[s[j]] >= have[s[j]] {
			count++
		}

		// 在满足条件的情况下尽可能的缩短窗口的大小
		for i <= j && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}

		// 计算当前的子串大小(也就是窗口大小)
		width := j - i - 1
		if count == len(t) && min > width {
			min = width
			res = s[i:j+1]
		}
		j++
	}

	return res
}

func main() {
	fmt.Println(minWindow(`ADOBECODEBANC`, `ABC`))
}
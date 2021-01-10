package main

import (
	"sort"
)

// 贪心
// 拿饼干给小孩子试试，如果满足ans++ 否则试下一组
func findContentChildren(greed, size []int) (ans int) {
	sort.Ints(greed)
	sort.Ints(size)
	g, s := 0, 0
	for g < len(greed) && s < len(size) {
		if greed[g] <= size[s] {
			g++
		}
		s++
	}

	return g
}

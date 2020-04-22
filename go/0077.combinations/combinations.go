package main

import (
	"fmt"
)

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
// 示例:
//
// 输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
// Related Topics 回溯算法

func combine(n int, k int) [][]int {
	var res [][]int

	if k <= 0 || n <= 0 || n < k {
		return res
	}

	var visit []int

	// 输入规模-递归层级-上一层产物-全局变量
	recursive(n, k, 1, visit, &res)
	return res
}

func recursive(n, k int, start int, visit []int, res *[][]int) {
	if len(visit) == k {
		tmp := make([]int, len(visit))
		copy(tmp, visit)

		*res = append(*res, tmp)
		return
	}

	for i := start; i <= n; i++ {
		visit = append(visit, i)

		recursive(n, k, i+1, visit, res)

		visit = visit[:len(visit)-1]

		fmt.Println(visit, cap(visit), len(visit))
	}

	return
}

//
//func combine2(n int, k int) [][]int {
//	var res [][]int
//
//	if k <= 0 || n <= 0 || n < k {
//		return res
//	}
//
//	var visit []int
//
//	// 输入规模-递归层级-上一层产物-全局变量
//	recursive2(n, k, 0, 1, visit, &res)
//	return res
//}
//
//func recursive2(n, k int, idx, begin int, visit []int, res *[][]int) {
//	if len(visit) == k {
//		temp := make([]int, k)
//		copy(temp, visit)
//		*res = append(*res, temp)
//		return
//	}
//
//	for i := begin; i <= n+1-k+idx; i++ {
//		visit[idx] = i
//		recursive2(n, k, idx+1, i+1, visit, res)
//		visit = visit[:len(visit)-1]
//	}
//}

func main() {
	fmt.Println(combine(4, 2))
	//fmt.Println(combine2(4, 2))
}

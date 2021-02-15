package main

import (
	"fmt"
)

//给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
//
// 例如，给出 n = 3，生成结果为：
//
// [
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]
//
// Related Topics 字符串 回溯算法

// 方法一：回溯法
//插入数量不超过n
//可以插入 ） 的前提是 ( 的数量大于 ）
func generateParenthesis(n int) []string {
	res := []string{}

	// 上一层产物-输入规模-递归层级-全局变量
	recursive(``, n, 0, 0, &res)
	return res
}

func recursive(cur string, n, left, right int, res *[]string) {
	if left == n && right == n {
		*res = append(*res, cur)
		return
	}

	// 左右判断条件只是剪枝作用
	if left < n {
		recursive(cur+`(`, n, left+1, right, res)
	}

	if left > right {
		recursive(cur+`)`, n, left, right+1, res)
	}
}

// 方法二:方法一的不同写法，但使用了byte更节省空间。方法一更直观。
func generateParenthesis2(n int) []string {
	res := []string{}
	cur := make([]byte, n*2)
	recursive2(cur, n, n, n, 0, &res)
	return res
}

func recursive2(cur []byte, n, left, right, idx int, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, string(cur))
	}

	// 可以分配左括号就一直往下分配
	if left > 0 {
		cur[idx] = '('
		recursive2(cur, n, left-1, right, idx+1, res)
	}

	//
	if left < right {
		cur[idx] = ')'
		recursive2(cur, n, left, right-1, idx+1, res)
	}

}

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}

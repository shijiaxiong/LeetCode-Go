package main

import (
	"fmt"
)

func subsets1(nums []int) [][]int {
	res := make([][]int, 1, 1024)

	for _, n := range nums {
		for _, r := range res {
			res = append(res, append([]int{n}, r...))
		}
	}

	return res
}

// 回溯算法
// 先序遍历。在回溯过程中记录深度。逐个考察数字，每个数都选或不选。
// https://leetcode-cn.com/problems/subsets/solution/hui-su-python-dai-ma-by-liweiwei1419/
func subsets(nums []int) [][]int {
	res := make([][]int, 0)

	// 这个循环主要是确定在哪层去记录信息
	// 会包含空子集所以i 从0开始
	for i := 0; i <= len(nums); i++ {
		cur := make([]int, 0, i)
		recursive(0, i, nums, cur, &res)
		fmt.Println(`----`)
	}

	return res
}

// depth 深度
// pos 当前所选nums中数字的索引
func recursive(pos, depth int, nums []int, cur []int, res *[][]int) {
	if len(cur) == depth {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := pos; i < len(nums); i++ {
		cur = append(cur, nums[i])
		fmt.Println(`-`, cur)

		recursive(i+1, depth, nums, cur, res)
		// 每次递归枚举的数字就减少
		cur = cur[:len(cur)-1]
	}
}


func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

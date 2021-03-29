package main

import (
	"sort"
)

//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用一次。
//
// 说明：
//
//
// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。
//
//
// 示例 1:
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
//所求解集为:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
//
//
// 示例 2:
//
// 输入: candidates = [2,5,2,1,2], target = 5,
//所求解集为:
//[
//  [1,2,2],
//  [5]
//]
// Related Topics 数组 回溯算法
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	sort.Ints(candidates)
	recursive(candidates, target, 0, cur, &res)

	return res
}

func recursive(candidates []int, target int, begin int, cur []int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := begin; i < len(candidates); i++ {
		// 剪枝 排序后 当前值大于target，则之后所有数都会大于target
		if candidates[i] > target {
			return
		}

		// 在进入递归之前都是当前层的逻辑
		if i > begin && candidates[i] == candidates[i-1] {
			continue
		}

		cur = append(cur, candidates[i])
		recursive(candidates, target-candidates[i], i+1, cur, res)
		cur = cur[:len(cur)-1]
	}
}

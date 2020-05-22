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
	sort.Ints(candidates)
	res := make([][]int, 0)
	tmp := make([]int, 0)

	recursive(candidates, target, 0, tmp, &res)
	return res
}

func recursive(candidates []int, target int, cur int, tmp []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, tmp)
		return
	}

	n := len(candidates)

	for i := cur; i < n; i++ {
		row := i
		// 跳过当前循环的重复项
		for i < n-1 && candidates[i] == candidates[i+1] {
			i++
		}

		// 当前项大于目标值返回
		if candidates[i] > target {
			return
		}

		new1 := make([]int, len(tmp))
		copy(new1, tmp)
		new1 = append(new1, candidates[i])
		recursive(candidates, target-candidates[i], row+1, new1, res)

	}
}

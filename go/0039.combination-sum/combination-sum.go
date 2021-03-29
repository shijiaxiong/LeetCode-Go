package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var cur []int

	recursive(candidates, target, 0, cur, &res)

	return res
}

func recursive(candidates []int, target int, begin int, cur []int, res *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	// 使用begin作为开始值 从每一层的第二个节点开始，都不能再搜索产生同一层结点已经使用过的candidates元素
	for i := begin; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		// 每个元素都可以重复使用，下一轮搜索的起点还是i
		recursive(candidates, target-candidates[i], i, cur, res)
		cur = cur[:len(cur)-1]
	}
}

func combinationSum1(candidates []int, target int) [][]int {
	var res [][]int

	var cur []int
	sort.Ints(candidates)
	recursive1(candidates, target, 0, cur, &res)

	return res
}

func recursive1(candidates []int, target int, begin int, cur []int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := begin; i < len(candidates); i++ {
		if target < candidates[i] {
			break
		}

		cur = append(cur, candidates[i])
		recursive1(candidates, target-candidates[i], i, cur, res)
		cur = cur[:len(cur)-1]
	}
}

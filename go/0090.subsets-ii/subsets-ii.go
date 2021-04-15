package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var cur []int
	sort.Ints(nums)

	recursive(nums, 0, cur, &res)

	return res
}

func recursive(nums []int, begin int, cur []int, res *[][]int) {

	tmp := make([]int, len(cur))
	copy(tmp, cur)
	*res = append(*res, tmp)

	for i := begin; i < len(nums); i++ {
		if i > begin && nums[i] == nums[i-1]{
			continue
		}

		cur = append(cur, nums[i])
		recursive(nums, i+1, cur, res)
		cur = cur[:len(cur)-1]
	}
}
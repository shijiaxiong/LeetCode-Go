package main

import (
	"fmt"
)

// 动态规划
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
// https://leetcode-cn.com/problems/subsets/solution/hui-su-python-dai-ma-by-liweiwei1419/
func subsets(nums []int) [][]int {
	res := make([][]int, 0)

	for i := 0; i <= len(nums); i++ {
		cur := make([]int, 0, i)
		recursive(0, i, nums, cur, &res)
	}

	return res
}

func recursive(pos, num int, nums []int, cur []int, res *[][]int) {
	if len(cur) == num {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := pos; i < len(nums); i++ {
		cur = append(cur, nums[i])
		recursive(i+1, num, nums, cur, res)
		cur = cur[:len(cur)-1]
	}
}


func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

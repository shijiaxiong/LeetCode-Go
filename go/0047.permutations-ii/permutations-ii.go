package main

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)

	cur := make([]int, length)
	visited := make([]bool, length)
	var res [][]int

	recursive(nums, length, 0, cur, visited, &res)

	return res
}

func recursive(nums []int, length int, depth int, cur []int, visited []bool, res *[][]int) {
	if length == depth {
		temp := make([]int, length)
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < length; i++ {
		if !visited[i] {
			// 访问过的相同数字跳过
			if i > 0 && nums[i] == nums[i-1] && visited[i-1] {
				continue
			}

			visited[i] = true
			cur[depth] = nums[i]

			recursive(nums, length, depth+1, cur, visited, res)

			visited[i] = false
		}
	}
}

func recursive0(nums []int, length int, depth int, path []int, visited []bool, res *[][]int) {
	if length == depth {
		temp := make([]int, length)
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	used := make(map[int]bool, length-depth)
	for i := 0; i < length; i++ {
		if !visited[i] && !used[nums[i]] {

			used[nums[i]] = true

			visited[i] = true
			path[depth] = nums[i]

			recursive0(nums, length, depth+1, path, visited, res)

			visited[i] = false
		}
	}
}

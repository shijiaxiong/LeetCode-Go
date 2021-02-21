package main

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
// 示例:
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
// Related Topics 回溯算法

// 方法一：回溯法
func permute(nums []int) [][]int {
	length := len(nums)

	path := make([]int, length)
	visited := make([]bool, length)
	var res [][]int

	recursive(nums, length, 0, path, visited, &res)

	return res
}

func recursive(nums []int, length int, depth int, path []int, visited []bool, res *[][]int) {
	if length == depth {
		temp := make([]int, length)
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := 0 ;i < length; i++ {
		if !visited[i] {
			visited[i] = true
			path[depth] = nums[i]

			recursive(nums, length, depth+1, path, visited, res)

			visited[i] = false
		}
	}
}

func main() {
	permute([]int{1, 2, 3})
}

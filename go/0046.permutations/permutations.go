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

// https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
// 方法一：回溯法
func permute(nums []int) [][]int {
	length := len(nums)

	// 因为cur的长度已经固定，所以recursive需要使用新的depth来表示递归层数
	cur := make([]int, length)
	// 记录访问状态
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

	for i := 0 ;i < length; i++ {
		if !visited[i] {
			visited[i] = true
			cur[depth] = nums[i]

			recursive(nums, length, depth+1, cur, visited, res)

			visited[i] = false
		}
	}
}

func main() {
	permute([]int{1, 2, 3})
}

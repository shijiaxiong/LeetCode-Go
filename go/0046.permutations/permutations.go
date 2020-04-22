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
	n := len(nums)

	vector := make([]int, n)

	taken := make([]bool, n)

	var ans [][]int

	makePermutation(0, n, nums, vector, taken, &ans)

	return ans
}

func makePermutation(cur, n int, nums, vector []int, taken []bool, ans *[][]int) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
	}

	for i := 0; i < n; i++ {
		if !taken[i] {
			taken[i] = true
			vector[cur] = nums[i]
			makePermutation(cur+1, n, nums, vector, taken, ans)
			taken[i] = false
		}
	}
}

func main() {
	permute([]int{1, 2, 3})
}

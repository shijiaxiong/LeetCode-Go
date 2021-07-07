package main

import "fmt"

// 直接递归
func minimumTotal(triangle [][]int) int {

	return dfs(triangle, 0, 0)
}

func dfs(triangle [][]int, i, j int) int {
	if i == len(triangle) {
		return 0
	}
	return min(dfs(triangle, i+1, j), dfs(triangle, i+1, j+1))+ triangle[i][j]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 记忆化递归
func minimumTotal0(triangle [][]int) int {
	mem := make([][]int, len(triangle))
	for i := range mem {
		mem[i] = make([]int, len(triangle[i]))
	}

	return dfs0(triangle, 0, 0, &mem)
}

func dfs0(triangle [][]int, i, j int, mem *[][]int) int {
	if i == len(triangle) {
		return 0
	}

	if (*mem)[i][j] != 0 {
		return (*mem)[i][j]
	}

	(*mem)[i][j] = min(dfs0(triangle, i+1, j, mem), dfs0(triangle, i+1, j+1,mem))+ triangle[i][j]

	return (*mem)[i][j]
}

// 动态规划
// 二维数组
// 时间复杂度:O(N^2)
// 空间复杂度:O(N^2)
func minimumTotal1(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle))
	}

	// 从左下角遍历
	for i := len(triangle) - 1; i>= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {//同一层数据
			//  base case
			if i == len(triangle) - 1 {
				dp[i][j] = triangle[i][j]
			} else {
				//  状态转移方程
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}

	return dp[0][0]
}

// 动态规划
// 一位数组
// 时间复杂度:O(N^2)
// 空间复杂度:O(N)
func minimumTotal2(triangle [][]int) int {
	// 多申请一个空间防止越界
	dp := make([]int, len(triangle) + 1)

	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

func main() {
	fmt.Println(minimumTotal0([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}

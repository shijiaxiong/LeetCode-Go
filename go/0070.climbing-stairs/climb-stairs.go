package problem0070

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
// 示例 1：
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//
// 示例 2：
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
//
// Related Topics 动态规划

// 解法一：暴力法
// 时间复杂度：O(2ⁿ) 树形递归总节点数  超时
// 空间复杂度：O(n) 树的深度
func climbStairs(n int) int {
	// i 定义了当前阶数，n 定义了目标阶数
	return recursive1(0, n)
}

func recursive1(i, n int) int {
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	return recursive1(i+1, n) + recursive1(i+2, n)
}

// 解法二：记忆化递归
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func climbStairs2(n int) int {
	mem := make(map[int]int)

	return recursive2(0, n, mem)
}

func recursive2(i, n int, mem map[int]int) int {
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	if mem[i] > 0 {
		return mem[i]
	}

	mem[i] = recursive2(i+1, n, mem) + recursive2(i+2, n, mem)
	return mem[i]
}

// 解法三：动态规划
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func climbStairs3(n int) int {
	if n < 2 {
		return 1
	}

	// 注意长度是n+11
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 解法四：斐波那契数列
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func climbStairs4(n int) int {
	if n < 2 {
		return 1
	}

	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		temp := first + second
		first = second
		second = temp
	}

	return second
}

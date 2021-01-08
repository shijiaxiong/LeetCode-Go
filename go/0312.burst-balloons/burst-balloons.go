package main

//动态规划
// 参考 https://leetcode-cn.com/problems/burst-balloons/solution/zhe-ge-cai-pu-zi-ji-zai-jia-ye-neng-zuo-guan-jian-/
// https://www.cnblogs.com/niuyourou/p/11964842.html
// https://mp.weixin.qq.com/s/I0yo0XZamm-jMpG-_B3G8g
func maxCoins(nums []int) int {
	n := len(nums)

	// 气球
	points := make([]int, n+2)
	points[0], points[n+1] = 1, 1

	for i := 1; i <= n; i++ {
		points[i] = nums[i-1]
	}

	dp := make([][]int, n+2)
	for k := range dp {
		dp[k] = make([]int, n+2)
	}

	// i 从下往上
	for i := n; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			// 区间内最后一个被戳破的气球
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+points[i]*points[k]*points[j]+dp[k][j])
			}
		}
	}
	return dp[0][n+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

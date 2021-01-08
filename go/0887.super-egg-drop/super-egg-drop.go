package main

// 动态规划
// 碎了 dp[i-1][j-1]
// 没碎 dp[i][j-1]
// 中间需要扔 1
// 因此 dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

	for j := 1; j <= N; j++ {
		for i := 1; i <= K; i++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
			//dp[i][j]大于楼层N的话，就可以返回楼层数j了
			if dp[i][j] >= N {
				return j
			}
		}
	}

	return N
}

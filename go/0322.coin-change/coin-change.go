package problem0322

import (
	"math"
	"sort"
)

// 动态规划
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if c <= i && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

// 递归 dfs + 剪枝
var res int
func coinChange1(coins []int, amount int) int {
	res = math.MaxInt32
	sort.Ints(coins)

	// 为了减少硬币数额，index 从最大值开始取
	dfs(coins, len(coins)-1, amount, 0)
	if res == math.MaxInt32 {
		return -1
	}

	return res
}

func dfs(coins []int, index int, amount int, cur int) {
	if index < 0 {
		return
	}

	for i := amount / coins[index]; i >= 0; i-- {
		// 匹配完对应的硬币后剩余金额
		remain := amount - i*coins[index]
		// 当前消耗硬币数
		curNum := cur + i

		// 剪枝1：已经凑足金额
		//大面值硬币需要更多小面值硬币替换，继续减少一枚或几枚大硬币搜索出来的答案【如果有】肯定不会更优。
		if remain == 0 {
			res = min(res, curNum)
			break
		}

		// 剪枝2：舍去 超过已知最小值情况
		// 当前的搜索结构amount>0,继续下去至少需要1个硬币，不会有更优解
		if curNum+1 >= res {
			break
		}

		dfs(coins, index-1, remain, curNum)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// 记忆化递归
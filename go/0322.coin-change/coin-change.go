package problem0322

import (
	"math"
	"sort"
)

// 动态规划
//https://leetcode-cn.com/problems/coin-change/solution/dong-tai-gui-hua-shi-yong-wan-quan-bei-bao-wen-ti-/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 理解 dp[0] = 0 的合理性，单独一枚硬币如果能够凑出面值，符合最优子结构
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		//注意：因为要比较的是最小值，这个不可能的值就得赋值成为一个最大值
		dp[i] = amount + 1
		for _, c := range coins {
			// 面值不超过要够凑的金额 amount
			// 当前要凑够i的硬币数(之前已经赋予不可能的最大值) 就是dp[i-c] + 1
			if c <= i && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	// 当前金额不能被给定硬币表示
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

// 递归 dfs + 剪枝
//https://leetcode-cn.com/problems/coin-change/solution/322-by-ikaruga/
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

//index-硬币下标；cur 当前层需要的硬币个数
func dfs(coins []int, index int, amount int, cur int) {
	if index < 0 {
		return
	}

	// 乘法加速，计算出最大能投入的数字 i
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
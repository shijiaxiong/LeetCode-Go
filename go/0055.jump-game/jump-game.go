package problem0055

// 贪心算法
func canJump0(nums []int) bool {
	// 能到达的最大范围
	k := 0
	for i := 0; i < len(nums); i++ {
		// 最大的累计步数都达不到当前的下标值，退出
		if i > k {
			return false
		}

		// 最大能覆盖的范围
		k = max(k, i+nums[i])
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 递归
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	return dfs(nums, 0)
}

func dfs(nums []int, index int) bool {
	// 递归终止条件
	if index >= len(nums)-1 {
		return true
	}

	// 根据nums[index] 往后走的每一个值进行递归
	for i := 1; i <= nums[index]; i++ {
		if dfs(nums, i+index) {
			return true
		}
	}

	return false
}

// 动态规划
// 拆分子元素 判断第二个元素能不能从第一个可达
func canJump1(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		//
		for j := i - 1; j >= 0; j-- {
		//for j := 0; j < i; j++ {
			//如果之前的j节点可达，并且从此节点可以到跳到i
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(nums)-1]
}

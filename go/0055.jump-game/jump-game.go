package problem0055

// 贪心算法
func canJump0(nums []int) bool {
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



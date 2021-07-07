package main

import (
	"fmt"
)

// 动态规划
// 因为存在负数 最小值可能会变最大值，所以需要保存最小值
func maxProduct(nums []int) int {
	res, imax, imin := nums[0],nums[0],nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			imax,imin = imin,imax
		}

		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])

		res = max(res, imax)
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxProduct([]int{-2,3,-4}))
}

//https://leetcode-cn.com/problems/maximum-product-subarray/solution/dong-tai-gui-hua-li-jie-wu-hou-xiao-xing-by-liweiw/
// 动态规划
// 时间复杂度:O(N)
// 空间复杂度:O(N)
func maxProduct1(nums []int) int {
	//  以 nums[i] 结尾的连续子数组的最值, 1 最大值，0最小值
	dp := make([][2]int, len(nums))
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] >= 0 {
			dp[i][0] = min(nums[i], dp[i-1][0] * nums[i])
			dp[i][1] = max(nums[i], dp[i-1][1] * nums[i])
		} else {
			dp[i][0] = min(nums[i], dp[i-1][1] * nums[i])
			dp[i][1] = max(nums[i], dp[i-1][0] * nums[i])
		}
	}

	res := dp[0][1]
	for i := 1; i < len(nums); i++ {
		res = max(res, dp[i][1])
	}

	return res
}

// 动态规划 降维

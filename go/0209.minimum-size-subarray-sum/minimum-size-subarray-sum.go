package main

import (
	"math"
)

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	sum := 0
	minLen := math.MaxInt32

	for right < len(nums) {
		// 滑动窗口右侧增大
		sum += nums[right]

		// 左侧减小，减小到不满足条件
		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++
	}

	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen
}
package main

import (
	"fmt"
	"sort"
)

// 排序 + 双指针
// 思路：先确定一个值i，然后在[i+1, len(nums)]两端使用双指针移动
// 时间复杂度：排序时间复杂度O(NlogN)+两层循环时间复杂度O(N2)
// 空间复杂度：排序空间复杂度O(logN) +O(N)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	ans := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums); i++ {
		// 剪枝去除重复元素, 当i>0时成立
		if i >0 && nums[i] == nums[i-1] {
			continue
		}

		start := i + 1 // 左指针
		end := len(nums) - 1 // 右指针

		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			fmt.Println(i, start, end)
			if abs(target-sum) < abs(target-ans) {
				ans = sum
			}

			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return ans
			}

		}
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

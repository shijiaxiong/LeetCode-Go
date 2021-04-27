package main

import "fmt"

func jump(nums []int) int {
	steps := 0
	start := 0 // 某一次起跳点范围开始的格子
	end := 1   // 某一次跳点范围结束的格子

	for end < len(nums) {
		maxPos := 0

		for i := start; i < end; i++ {
			// 能跳到最远的距离
			maxPos = max(maxPos, i+nums[i])
		}

		start = end      // 下一次起跳点范围开始的格子
		end = maxPos + 1 // 下一次跳点范围结束的格子
		steps++            // 跳跃次数
	}

	return steps
}

// 优化
func jump0(nums []int) int {
	end := 0
	maxPosition := 0
	steps := 0

	// 题目中确保一定会达到终点，当到达终点的时候已经不需要再跳一次
	for i := 0; i < len(nums)-1; i++ {
		//找能跳的最远的
		maxPosition = max(maxPosition, nums[i]+i)

		//遇到边界，就更新边界，并且步数加一
		// 边界由上一次跳跃确定
		if i == end {
			end = maxPosition // 找到下一次的边界
			steps++
		}
	}

	return steps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//[3,4,3,2,5,4,3]
	fmt.Println(jump([]int{3, 4, 3, 2, 5, 4, 3}))
}

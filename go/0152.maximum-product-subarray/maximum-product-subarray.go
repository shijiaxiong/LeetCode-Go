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
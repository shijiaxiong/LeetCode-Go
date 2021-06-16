package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/powx-n/solution/powx-n-by-leetcode-solution/
// 递归
// 时间复杂度：O(logn)
// 空间复杂度：O(logn)
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	// 递归指数减少一半
	y := quickMul(x, n/2)

	if n%2 == 0 {
		return y * y
	}

	// 奇数次
	return y * y * x
}

// 迭代
// 时间复杂度：O(logn)
// 空间复杂度：1
func quickMul0(x float64, N int) float64 {
	ans := 1.0
	// 在对 N 进行二进制拆分的同时计算答案
	for N > 0 {
		if N%2 == 1 {
			// 如果 N 二进制表示的最低位为 1，那么需要计入贡献
			ans *= x
		}

		// 将贡献不断地平方
		x *= x
		// 舍弃 N 二进制表示的最低位，这样我们每次只要判断最低位即可
		N /= 2
	}
	return ans
}

// 直接迭代
func myPow0(x float64, n int) float64 {
	if n < 0 {
		x = 1.0 / x
		n = -n
	}

	ans := 1.0

	for n > 0 {
		if n%2 == 1 {

			ans *= x
		}
		fmt.Println(n)

		x *= x
		n >>= 1
	}

	return ans
}

func main() {
	myPow0(10.0, 4)
}
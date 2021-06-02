package proble0069

import (
	"fmt"
)

// 牛顿法
func mySqrt0(x int) int {
	a := x
	for a*a > x {
		a = (a + x/a) / 2
	}
	return a
}

// 二分法
func mySqrt(x int) int {
	left := 1
	right := x/2 + 1

	for left <= right {

		mid := left + (right-left)/2

		fmt.Println(left, mid , right)

		if mid == x/mid {
			return mid
		} else if mid > x/mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 可能存在left=right+1
	return right
}

// 二分法
// https://leetcode-cn.com/problems/sqrtx/solution/er-fen-cha-zhao-niu-dun-fa-python-dai-ma-by-liweiw/
func mySqrt1(x int) int {
	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	left := 1
	right := x/2

	// 注意到题目中给出的例 2，小数部分将被舍去。我们就知道了，如果一个数 aa 的平方大于 xx ，那么 aa 一定不是 xx 的平方根。我们下一轮需要在 [0..a - 1][0..a−1] 区间里继续查找 xx 的平方根。
	for left < right {
		mid := left + (right - left + 1)/2

		if mid > x / mid {
			right= mid-1
		}  else {
			left = mid
		}
	}

	return left
}
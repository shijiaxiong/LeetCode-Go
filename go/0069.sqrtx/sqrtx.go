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

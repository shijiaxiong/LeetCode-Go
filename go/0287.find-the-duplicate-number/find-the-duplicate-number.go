package main

// 二分法
// 时间复杂度：O(NlogN)
// 空间复杂度：O(1）
func findDuplicate(nums []int) int {
	//  left right 是数据范围
	left := 1
	right := len(nums) - 1

	for left < right {
		//  向下取整，如果数组长度是奇数个,mid就是他们的中间值，小于等于mid的数
		mid := left + (right - left) / 2

		cnt := 0
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}

		if cnt <= mid {
			left  = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// 快慢指针
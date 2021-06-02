package main

// 转换为寻找第一个大于等于目标元素的下标位置
func searchInsert(nums []int, target int) int {
	if nums[len(nums) - 1] < target {
		return len(nums)
	}

	// 区间 [nums[left...right]
	left := 0
	right := len(nums) - 1

	for left < right {
		// 向下取整，mid在左边范围
		mid := left + (right - left) / 2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}

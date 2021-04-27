package main

func searchInsert(nums []int, target int) int {
	if nums[len(nums) - 1] < target {
		return len(nums)
	}

	// 区间 [nums[left...right]
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right - left) / 2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}

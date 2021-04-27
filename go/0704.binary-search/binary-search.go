package main

func search(nums []int, target int) int {
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

	if nums[left] == target {
		return left
	}

	return -1
}

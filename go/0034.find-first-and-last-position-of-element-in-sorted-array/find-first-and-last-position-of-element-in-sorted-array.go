package main

// 直接寻找first last的二分查找
// 注意mid值
func searchRange2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first := findFirst(nums, target)
	last := findLast(nums, target)
	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low < high {
		// mid 向下取整
		mid := low + (high-low)/2
		//fmt.Println(`first`,low, high, mid)
		switch {
		case nums[mid] < target:
			low = mid + 1
		case nums[mid] > target:
			high = mid - 1
		default:
			high = mid
		}
	}

	if nums[low] == target {
		return low
	}

	return -1
}

func findLast(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low < high {
		// mid 向上取整
		mid := low + (high-low+1)/2
		//fmt.Println(`last`, low, high, mid)
		switch {
		case nums[mid] < target:
			low = mid + 1
		case nums[mid] > target:
			high = mid - 1
		default:
			low = mid
		}
	}

	return low
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left :=0
	right := len(nums) - 1

	for left < right {
		mid := left + (right - left + 1)/ 2

		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}

	start := -1
	end := -1
	if nums[left] == target {
		start = left
		end = left
		for start > 0 && nums[start - 1] == nums[start] {
			start--
		}

		// 这个循环可以不写 因为向上取整之后得到的值就是右边界
		for end < len(nums) - 1 && nums[end] == nums[end + 1] {
			end++
		}
		return []int{start, end}
	}

	return []int{start, end}
}
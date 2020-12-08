package main

import "fmt"

// 二分查找法
// 时间复杂度：N*O(logN)
// 空间复杂度：O(1)
func searchRange(nums []int, target int) []int {
	index := search(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}

	first := index
	for {
		f := search(nums[:first], target)
		if f == -1 {
			break
		}

		first = f
	}

	last := index
	for {
		l := search(nums[last+1:], target)
		if l == -1 {
			break
		}

		// 新的数组从0开始计算，所以相对于first需要再加1
		last += l + 1
	}

	return []int{first, last}
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var medium int

	for low <= high {
		medium = (low + high) / 2
		switch {
		case nums[medium] < target:
			low = medium + 1
		case nums[medium] > target:
			high = medium - 1
		default:
			return medium
		}
	}

	return -1
}

func main() {
	res := searchRange2([]int{5, 7, 7, 8, 8, 8, 10}, 8)
	fmt.Println(res)
}

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

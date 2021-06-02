package main

import "fmt"

// 二分法
// 无限分割数组，永远都在有序的子数组中去查找target
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)>>1

		fmt.Println(low, mid, high)

		if nums[mid] == target {
			return mid
		}

		// 小于等于原因：
		//当剩下两个数字的时候,start和mid相同
		if nums[low] < nums[mid] {
			// 要找的数在前半部分
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// 要找的数在前半部分
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func main(){
	fmt.Println(search([]int{3,1}, 1))
}

// 思路：只判断有序的那部分区间，其他情况放在else里
//      先判断哪部分有序，找到有序区间。然后
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/er-fen-fa-python-dai-ma-java-dai-ma-by-liweiwei141/
func search0(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		// 根据分支的逻辑将中间数改成上取整
		mid := left + (right - left + 1) / 2

		//  右区间严格有序
		if nums[mid] < nums[right] {
			// 此时 [mid..right] 有序
			if nums[mid] <= target && target <= nums[right] {
				// 如果 target 的值落在这个区间里，下一轮搜索区间是 [mid..right]，此时设置 left = mid;
				left = mid
			} else {
				// 否则，下一轮搜索区间是 [left..mid - 1]，此时设置 right = mid - 1;
				right = mid - 1
			}
		} else {
			// 此时 nums[mid] >= nums[right]，注意此时 mid 可能与 right 重合
			// 数组前半部分有序，即 [left..mid] 有序，为了与上一个分支的逻辑一致，认为 [left..mid - 1] 有序
			if nums[left] <= target && target <= nums[mid - 1] {
				right = mid - 1
			} else {
				// 否则，下一轮搜索区间是 [mid..right]，此时设置 left = mid;
				left = mid
			}
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}
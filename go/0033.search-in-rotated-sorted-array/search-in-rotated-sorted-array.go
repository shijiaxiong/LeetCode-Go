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
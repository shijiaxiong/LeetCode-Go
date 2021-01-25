package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	// 排序后，可以按规律查找
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		// 排序后的数字，如果当前数字大于0，则三数之和一定大于0
		if nums[i] > 0 {
			break
		}

		// 避免添加重复的结果
		// i>0 是为了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				// 较小的 l 需要变大
				l++
			case s > 0:
				// 较大的 r 需要变小
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				//l, r = next(nums, l, r)
				for l <r && nums[l] == nums[l+1] {
					l++
				}
				for l<r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}
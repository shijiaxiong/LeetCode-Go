package main

//给你一个整数数组 nums，请你将该数组升序排列。
//
//
//
//
//
//
// 示例 1：
//
// 输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
//
//
// 示例 2：
//
// 输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 50000
// -50000 <= nums[i] <= 50000
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	quickSort(nums)
	return nums
}

// 方法一： 快排
func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	reference := nums[0]
	head, tail := 0, len(nums)-1
	i := 1

	for head < tail {
		if nums[i] < reference {
			nums[head], nums[i] = nums[i], nums[head]
			head++
			i++
		} else {
			nums[tail], nums[i] = nums[i], nums[tail]
			tail--
		}
	}

	quickSort(nums[:head])
	quickSort(nums[head+1:])
	return
}

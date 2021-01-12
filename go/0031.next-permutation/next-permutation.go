package main

//https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1

	//	从后向前寻找第一个相邻的元素对(i,j)满足A[i]<A[j]
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		//	在[j,end)从后向前查找第一个满足A[i]<A[k]的k
		for nums[i] >= nums[k] {
			k--
		}

		// 交换A[i] A[k]的位置
		nums[i], nums[k] = nums[k], nums[i]
	}

	// 翻转 A[j:end]
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

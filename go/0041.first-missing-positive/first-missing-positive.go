package main

// 哈希表
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func firstMissingPositive(nums []int) int {
	length := len(nums)

	for i := 0; i < length; i++ {
		//如果nums[i]在[1,len(nums)]之间，将nums[i]放到对应下标的位置
		for nums[i] > 0 && nums[i] <= length && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// 再次遍历数组
	for i := 0; i < length; i++ {
		if nums[i] != i+1 { //当前位置是缺失的第一个正数
			return i + 1
		}
	}

	return length + 1
}

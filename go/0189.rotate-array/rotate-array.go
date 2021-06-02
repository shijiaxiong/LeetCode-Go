package main

// 多次旋转替换数组
func rotate1(nums []int, k int) {
	for i := 0; i < k; i++ {
		traget := nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			nums[j], traget = traget, nums[j]
		}
	}
}

// 旋转两次数组。
func rotate(nums []int, k int) {
	if k > len(nums) {
		k = k%len(nums)
	}

	if k == 0 || k == len(nums) {
		return
	}

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

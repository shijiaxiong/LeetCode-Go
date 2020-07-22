package main

func moveZeroes(nums []int) {
	l := len(nums)
	i, j := 0, 0

	// i,j 两个指针在数组内移动
	// j 先行，寻找非0元素，找到后与i的值调换位置。i移动到下一个位置
	for j < l {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j],nums[i]
			i++
		}

		j++
	}

	// fmt.Println(nums)
}

package main

import "fmt"

func findDuplicates(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		// 数字出现过一次就会为负数所以需要取绝对值
		// 如果数字和下标要一一对应，nums[i]在abs(nums[i]) - 1的位置
		index := abs(nums[i]) - 1

		// 如果这个位置的值已经是负值， 说明已经出现过一次，加入res
		if nums[index] < 0 {
			res = append(res, abs(nums[i]))
		}

		// 将对应下标处的值记为负数
		nums[index] = -nums[index]
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

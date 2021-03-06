package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0

	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
		fmt.Println(nums)
	}

	return i + 1
}

func main()  {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}

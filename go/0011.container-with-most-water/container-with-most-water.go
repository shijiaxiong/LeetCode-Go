package main

import (
	"fmt"
)

// 双指针
func maxArea(height []int) int {
	i, j, max := 0, len(height)-1, 0

	for i < j {
		if height[i] < height[j] {
			max = getMax(max, (j-i) * height[i])
			i++
		} else {
			max = getMax(max, (j-i) * height[j])
			j--
		}
	}

	return max
}

func getMax(i, j int) int {
	if i <= j {
		return j
	}
	return i
}

func main() {
	res := maxArea([]int{1,8,6,2,5,4,8,3,7})
	fmt.Println(res)
}

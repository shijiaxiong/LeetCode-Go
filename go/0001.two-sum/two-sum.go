package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))

	for k, v := range nums {
		fmt.Println(index)

		if m, ok := index[target-v]; ok {
			return []int{m, k}
		}

		index[v] = k
	}

	fmt.Println(index)

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15},9))
}


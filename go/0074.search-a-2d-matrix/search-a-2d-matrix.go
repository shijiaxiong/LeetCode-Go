package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		//  寻找目标值会在哪一行
		if matrix[i][0] <= target && target <= matrix[i][len(matrix[i])  - 1] {
			left := 0
			right := len(matrix[i]) - 1
			for left < right {
				mid := left + (right - left + 1) / 2
				if matrix[i][mid] > target {
					right = mid - 1
				} else {
					left = mid
				}
			}

			if matrix[i][left] == target {
				return true
			}

			return false
		}
	}

	return false
}

// 把矩阵退化成一位数组
func searchMatrix1(matrix [][]int, target int) bool {
	left := 0
	right := len(matrix) * len(matrix[0]) - 1
	m := len(matrix[0])

	for left < right {
		//  向上取整[left, mid-1] [mid, right]
		mid := left + (right - left + 1) / 2
		if matrix[mid/m][mid%m] > target {// mid在右区间，mid值>target的时候，目标值一定在左区间
			right = mid - 1
		} else {
			left = mid
		}
	}

	if matrix[left/m][left%m] == target {
		return true
	}

	return false
}

func main (){
	fmt.Println(searchMatrix([][]int{{1,3,5,7}, {10,11,16,20},{23,30,34,50}},13))
}
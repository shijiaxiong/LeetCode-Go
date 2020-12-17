package main

import (
	"fmt"
)

func searchMatrix0(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) == 0{
			return false
		}
		// 判断目标值在哪一行
		if matrix[i][0] <= target && matrix[i][len(matrix[i])-1] >= target {
			l := 0
			h := len(matrix[i])-1
			for l <= h {
				mid := l+ (h - l) / 2
				if matrix[i][mid] > target {
					h = mid-1
				} else if matrix[i][mid] < target {
					l = mid+1
				} else {
					return true
				}
			}
		}
	}

	return false
}

// 把矩阵退化成一位数组
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m := len(matrix[0])
	low := 0
	high := len(matrix[0])*len(matrix)-1

	for low <= high {
		mid := low + (high - low) >> 1
		// 找到虚数组的序号
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] >target {
			high = mid -1
		} else {
			low = mid +1
		}
	}
	return false
}

func main (){
	fmt.Println(searchMatrix([][]int{{1,3,5,7}, {10,11,16,20},{23,30,34,50}},13))
}
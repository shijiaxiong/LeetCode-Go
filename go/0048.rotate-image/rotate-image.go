package main

import "fmt"

// 时间复杂度:O(N2)
// 空间复杂度:O(1)
//  先以对角线(左上<->右下)为轴进行翻转，再对每行进行左右翻转
func rotate(matrix [][]int) {
	n := len(matrix)

	//以对角线(左上<->右下)为轴进行翻转
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i],matrix[i][j]
		}
	}
	fmt.Println(matrix)

	//	 以每行进行左右翻转
	mid := n >> 1
	for i := 0; i < n; i++ {
		for j := 0; j < mid; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

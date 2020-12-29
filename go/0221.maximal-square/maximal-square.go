package main

//public int maximalSquare(char[][] matrix) {
//    // base condition
//    if (matrix == null || matrix.length < 1 || matrix[0].length < 1) return 0;
//
//    int height = matrix.length;
//    int width = matrix[0].length;
//    int maxSide = 0;
//
//    // 相当于已经预处理新增第一行、第一列均为0
//    int[][] dp = new int[height + 1][width + 1];
//
//    for (int row = 0; row < height; row++) {
//        for (int col = 0; col < width; col++) {
//            if (matrix[row][col] == '1') {
//                dp[row + 1][col + 1] = Math.min(Math.min(dp[row + 1][col], dp[row][col + 1]), dp[row][col]) + 1;
//                maxSide = Math.max(maxSide, dp[row + 1][col + 1]);
//            }
//        }
//    }
//    return maxSide * maxSide;
//}
//
func maximalSquare(matrix [][]byte) int {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return 0
	}

	height := len(matrix)
	width := len(matrix[0])
	maxSide := 0

	dp := make([][]int, height+1)
	for i := range dp {
		dp[i] = make([]int, width+1)
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if matrix[row][col] == '1' {
				dp[row+1][col+1] = min(min(dp[row+1][col], dp[row][col+1]), dp[row][col]) + 1
				maxSide = max(maxSide, dp[row+1][col+1])
			}
		}
	}
	return maxSide * maxSide
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

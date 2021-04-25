package main

func islandPerimeter(grid [][]int) int {
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				// 题目限制只有一个岛屿，计算一个即可
				res = dfs(&grid, i, j)
			}
		}
	}

	return res
}

func dfs(grid *[][]int, i, j int) int {
	// 整个网格的边界，算一条边
	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) {
		return 1
	}

	// 连接海洋的地方，算一条边
	if (*grid)[i][j] == 0 {
		return 1
	}

	//  函数因为「当前格子是已遍历的陆地格子」返回，和周长没关系
	if (*grid)[i][j] != 1 {
		return 0
	}

	(*grid)[i][j] = 2

	return dfs(grid, i+1, j) +
		dfs(grid, i-1, j) +
		dfs(grid, i, j-1) +
		dfs(grid, i, j+1)
}

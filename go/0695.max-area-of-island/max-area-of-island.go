package problem0695

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	for i := range grid {
		for j := range grid[i] {
			maxArea = max(maxArea, dfs(&grid, i, j))
		}
	}

	return maxArea
}

func dfs(grid *[][]int, r, c int) int {
	//如果坐标（r,c）超出了网格范围，直接返回
	if !((0 <= r && r < len(*grid)) && 0 <= c && c < len((*grid)[0])) {
		return 0
	}

	//如果这个格子不是岛屿，直接返回
	if (*grid)[r][c] != 1 {
		return 0
	}

	// 将格子标记为【已遍历过】
	(*grid)[r][c] = 2

	// 访问上下左右的节点
	return 1 + dfs(grid, r-1, c) + dfs(grid, r+1, c) + dfs(grid, r, c-1) + dfs(grid, r, c+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

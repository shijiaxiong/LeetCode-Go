package problem0200

// 解法一：dfs
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(&grid, i,j)
			}
		}
	}

	return count
}

func dfs(grid *[][]byte, r, c int) {
	//如果坐标（r,c）超出了网格范围，直接返回
	if !isArea(grid, r, c) {
		return
	}

	//如果这个格子不是岛屿，直接返回
	if (*grid)[r][c] != '1' {
		return
	}

	// 将格子标记为【已遍历过】
	(*grid)[r][c] = '2'

	// 访问上下左右的节点
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

//判断坐标（r,c）是否在网格中
func isArea(grid *[][]byte, r, c int) bool {
	return (0 <= r && r < len(*grid)) && 0 <= c && c < len((*grid)[0])
}

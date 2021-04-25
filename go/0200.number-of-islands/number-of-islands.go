package problem0200

// 解法一：dfs
// https://leetcode-cn.com/problems/number-of-islands/solution/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 岛屿中的点  并且没有被访问过 就进行深度优先遍历
			if grid[i][j] == '1' {
				count++
				dfs(&grid, i, j)
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
	// 避免重复
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

// 并查集
// 联通量的个数减去空地的个数就是岛屿的数量
func numIslands0(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	ufind := NewUnionFind(rows * cols)
	space := 0

	// 将二维表转成一维点
	node := func(i, j int) int {
		return i*cols + j
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				space++
			} else {
				if i > 0 && grid[i-1][j] == '1' {
					ufind.Union(node(i, j), node(i-1, j))
				}
				if i < rows-1 && grid[i+1][j] == '1' {
					ufind.Union(node(i, j), node(i+1, j))
				}
				if j > 0 && grid[i][j-1] == '1' {
					ufind.Union(node(i, j), node(i, j-1))
				}
				if j < cols-1 && grid[i][j+1] == '1' {
					ufind.Union(node(i, j), node(i, j+1))
				}
			}
		}
	}

	return ufind.count-space
}

type UnionFind struct {
	count int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &UnionFind{
		parent: parent,
		count: n,
	}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
		u.count--
	}
}

func (u *UnionFind) Find(i int) int {
	node := i
	for u.parent[node] != node {
		u.parent[node] = u.parent[u.parent[node]]
		node = u.parent[node]
	}

	return node
}

package main

// 问题转化成寻找和边界联通的O
// dfs 递归
func solve0(board [][]byte) {
	if len(board) <= 0 {
		return
	}

	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 从边缘搜索
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}

	return
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == 'X' || board[i][j] == '#' {
		return
	}

	board[i][j] = '#'
	dfs(board, i-1, j) // 上
	dfs(board, i+1, j) // 下
	dfs(board, i, j-1) // 左
	dfs(board, i, j+1) // 右
}

// 并查集
func solve(board [][]byte) {
	if len(board) <= 0 {
		return
	}

	rows, cols := len(board), len(board[0])
	// ufind是虚拟节点，边界上O的父节点都是这个虚拟节点
	dummy, ufind := rows*cols, NewUnionFind(rows*cols+1)

	// 将二维表转成一维点
	node := func(i, j int) int {
		return i*cols + j
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' { // 遇到O进行并查集操作合并
				//	 边界上的O 与dummy进行合并
				if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
					ufind.Union(node(i, j), dummy)
				} else {
					if i > 0 && board[i-1][j] == 'O' {
						ufind.Union(node(i, j), node(i-1, j))
					}
					if i < rows-1 && board[i+1][j] == 'O' {
						ufind.Union(node(i, j), node(i+1, j))
					}
					if j > 0 && board[i][j-1] == 'O' {
						ufind.Union(node(i, j), node(i, j-1))
					}
					if j < cols-1 && board[i][j+1] == 'O' {
						ufind.Union(node(i, j), node(i, j+1))
					}
				}
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			//if ufind.IsConnected(node(i,j), dummy) {
			//	board[i][j] = 'O'
			//} else {
			//	board[i][j] = 'X'
			//}
			//
			if board[i][j] == 'O' && !ufind.IsConnected(node(i,j), dummy) {
				board[i][j] = 'X'
			}
		}
	}

}

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
	}
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
	}
}

func (u *UnionFind) Find(i int) int {
	node := i
	// 当前节点的父节点 指向父节点的父节点
	// 保证联通区域的父节点只有一个
	for u.parent[node] != node {
		u.parent[node]= u.parent[u.parent[node]]
		node = u.parent[node]
	}

	return node
}

func (u *UnionFind) IsConnected(i, j int) bool {
	return u.Find(i) == u.Find(j)
}

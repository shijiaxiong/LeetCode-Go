package main

func solveNQueens(n int) [][]string {
	if n < 1 {
		return [][]string{}
	}

	cols := make([]bool, n)
	pie := make([]bool, 2*n)
	na := make([]bool, 2*n)

	res := [][]string{}
	board := make([]string, n) // 运算中间结果,指定数组的长度可以不用再次传递n
	dfs(cols, pie, na, board, 0, &res)
	return res
}

func dfs(cols, pie, na []bool, board []string, r int, res *[][]string) {
	n := len(board)

	// terminator
	if n == r {
		tmp := make([]string, len(board))
		copy(tmp, board)
		*res = append(*res, tmp)
		return
	}

	// recursive
	for c := 0; c < n; c++ {
		// na满足 r-c=常数 ，加n是为了不出现负数
		na1 := r - c + n

		// pie满足 r+c=常数
		// 原答案：pie1 = 2*n - r - c - 1
		pie1 := r + c
		if !cols[c] && !pie[pie1] && !na[na1] {

			b := make([]byte, n)
			for i := range b {
				b[i] = '.'
			}
			b[c] = 'Q'
			board[r] = string(b)

			cols[c], pie[pie1], na[na1] = true, true, true

			dfs(cols, pie, na, board, r+1, res)

			//revert the current level states
			cols[c], pie[pie1], na[na1] = false, false, false
		}
	}
}


package problem0429

type Node struct {
	Val      int
	Children []*Node
}

var res [][]int

// 递归
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res = [][]int{}
	dfs(root, 0)
	return res
}

func dfs(root *Node, level int) {
	if root == nil {
		return
	}

	// 初始化当前层
	if len(res) == level {
		res = append(res, []int{})
	}

	res[level] = append(res[level], root.Val)
	for _, n := range root.Children {
		dfs(n, level+1)
	}
}

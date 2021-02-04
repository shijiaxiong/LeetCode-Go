package main

type Node struct {
	Val      int
	Children []*Node
}

// 递归
func preorder(root *Node) []int {
	var res  []int

	dfs(root, &res)

	return res
}

func dfs(root *Node, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)

	for _,children :=range root.Children {
		dfs(children, res)
	}
}

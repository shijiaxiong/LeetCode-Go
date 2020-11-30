package problem0590

type Node struct {
	Val int
	Children []*Node
}

var res []int

func postorder(root *Node) []int {
	dfs(root)
	return res
}

func dfs(root *Node)  {
	if root != nil {
		for _, c := range root.Children{
			dfs(c)
		}
		res = append(res, root.Val)
	}
}
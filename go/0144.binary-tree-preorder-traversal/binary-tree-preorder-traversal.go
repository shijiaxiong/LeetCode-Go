package problem0144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
// 解法一:递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int

	res = append(res, root.Val)

	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)

	return res
}

// 解法二:递归
func preorderTraversal1(root *TreeNode) []int {
	var res []int
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	dfs(root.Left, res)
	dfs(root.Right, res)
}

// 解法三:迭代

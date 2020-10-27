package problem0094

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)//切片打散参入

	return res
}

// 解法二:递归
func inorderTraversal1(root *TreeNode) []int {
	var res []int
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, res)
	*res = append(*res, root.Val)
	dfs(root.Right, res)
}

// 解法三:迭代

package problem0094

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...) //切片打散参入

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
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root) //入栈
			root = root.Left            //移至最左
		}

		root = stack[len(stack)-1] // 栈顶
		stack = stack[:len(stack)-1]
		res = append(res, root.Val) // 中序输出
		root = root.Right           // 右节点会进入下次循环，如果 =nil，继续出栈
	}

	return res
}

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
func preorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for len(stack) > 0 || root != nil { // root!=nil 只会在第一次判断root时使用 放到最后
		for root != nil {
			res = append(res, root.Val) // 前序输出
			stack = append(stack, root)  //
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

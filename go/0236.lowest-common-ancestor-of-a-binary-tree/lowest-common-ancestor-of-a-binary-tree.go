package main

type TreeNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// terminator-终结状态
	if root == nil || root == q || root ==p {
		return root
	}

	//	process current logic

	// drill down-递归
	l := lowestCommonAncestor(root.Left, p,q)
	r := lowestCommonAncestor(root.Right, p,q)

	// restore current status-处理当前层数据
	// l == nil,意味着p,q都不在root的左子树中
	// r == nil,意味着p,q都不在root的右子树中
	if l != nil && r != nil {
		return root
	}

	if l ==nil {
		return r
	}

	return l
}
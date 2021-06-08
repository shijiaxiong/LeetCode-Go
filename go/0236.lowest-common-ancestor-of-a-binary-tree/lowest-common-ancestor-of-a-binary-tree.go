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

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/236-er-cha-shu-de-zui-jin-gong-gong-zu-xian-hou-xu/
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	// 递归终止条件
	// 到达叶子结点，再往下就是空；root 等于 p 或者q 直接返回root
	if root == nil || root == q || root == p {
		return root
	}

	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)

	// root的左右子树中都不含有p q 返回空
	if left == nil && right == nil {
		return nil
	}

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	// left right同时不为空，p,q在异侧
	return root
}
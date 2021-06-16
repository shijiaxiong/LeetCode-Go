package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--22/
//
func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeHelper(preorder, 0, len(preorder), inorder, 0, len(inorder))
}

func buildTreeHelper(preorder []int, pStart, pEnd int, inorder []int, iStart, iEnd int) *TreeNode {
	// preorder 为空 直接返回null
	if pStart == pEnd {
		return nil
	}

	// 根节点值
	rootVal := preorder[pStart]
	root := &TreeNode{
		Val : rootVal,
	}

	// 中序遍历中找到跟节点的位置
	iRootIndex := 0
	for i := iStart; i < iEnd; i++ {
		if rootVal == inorder[i] {
			iRootIndex = i
			break
		}
	}
	// 根节点左边就是左子树的部分
	leftNum := iRootIndex - iStart

	//  递归构造左子树
	root.Left = buildTreeHelper(preorder, pStart + 1, pStart + leftNum + 1, inorder, iStart, iRootIndex)
	//  递归构造右子树
	root.Right = buildTreeHelper(preorder, pStart + leftNum + 1, pEnd, inorder, iRootIndex + 1, iEnd)
	return root
}
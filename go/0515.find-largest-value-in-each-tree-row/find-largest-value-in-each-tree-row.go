package problem0515

import (
	"math"
)

//type TreeNode utils.TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		max := math.MinInt64
		maxLen := len(queue) // 防止queue变化导致循环次数发生变化
		for i := 0; i < maxLen; i++ {
			node := queue[0]
			if max < node.Val {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}

		res = append(res, max)
	}

	return res
}


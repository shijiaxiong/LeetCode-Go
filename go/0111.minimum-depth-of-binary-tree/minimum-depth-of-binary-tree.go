package problem0111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先
func minDepth(root *TreeNode) int {
	if root ==nil {
		return 0
	}

	if root.Left != nil && root.Right != nil {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	} else if root.Left != nil {
		return 1+minDepth(root.Left)
	} else if root.Right != nil {
		return 1+minDepth(root.Right)
	} else {
		return 1
	}

}

func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}

// 广度优先
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// q存储的是下层节点
	q := []*TreeNode{root}
	depth := 1

	for len(q) != 0 {
		size := len(q)
		for i :=0 ;i < size;i++ {
			cur := q[0]
			q = q[1:]

			if cur.Left == nil  && cur.Right == nil {
				return depth
			}

			if cur.Left != nil {
				q = append(q, cur.Left)
			}

			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}

	return depth
}
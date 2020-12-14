package problem0102

import (
	"github.com/shijiaxiong/LeetCode-Go/go/utils"
)

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func levelOrder(root *TreeNode) [][]int {
//	res := [][]int{}
//	var dfs func(*TreeNode, int)
//
//	dfs = func(root *TreeNode, level int) {
//		if root == nil {
//			return
//		}
//
//		// 出现了新的 level
//		if level >= len(res) {
//			res = append(res, []int{})
//		}
//		res[level] = append(res[level], root.Val)
//
//		dfs(root.Left, level+1)
//		dfs(root.Right, level+1)
//	}
//
//	dfs(root, 0)
//
//	return res
//}
//
//// dfs的常规写法
//var res [][]int
//
//func levelOrder1(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//	res = make([][]int, 0)
//	dfs(root, 0)
//	return res
//}
//
//func dfs(root *TreeNode, level int) {
//	if root == nil {
//		return
//	}
//	if level == len(res) {
//		res = append(res, []int{})
//	}
//	res[level] = append(res[level], root.Val)
//	dfs(root.Left, level+1)
//	dfs(root.Right, level+1)
//}
//
//// BFS 广度优先
//// 使用较多queue函数
//func levelOrder2(root *TreeNode) [][]int {
//	var result [][]int
//	if root == nil {
//		return result
//	}
//
//	// 定义一个双向队列
//	queue := list.New()
//	// 头部插入根节点
//	queue.PushFront(root)
//	// 进行广度搜索
//	for queue.Len() > 0 {
//		var current []int
//		listLength := queue.Len()
//		for i := 0; i < listLength; i++ {
//			//消耗尾部
//			// queue.Remove(queue.Back()).(*TreeNode)：移除最后一个元素并将其转化为TreeNode类型
//			node := queue.Remove(queue.Back()).(*TreeNode)
//			current = append(current, node.Val)
//			if node.Left != nil {
//				//插入头部
//				queue.PushFront(node.Left)
//			}
//			if node.Right != nil {
//				queue.PushFront(node.Right)
//			}
//		}
//		result = append(result, current)
//	}
//	return result
//}

// BFS 广度优先
func levelOrder3(root *utils.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue []*utils.TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		var tmp []int

		maxLen := len(queue) // 防止queue变化导致循环次数发生变化
		for i := 0; i < maxLen; i++ {
			node := queue[0]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}

		res = append(res, tmp)
	}

	return res
}

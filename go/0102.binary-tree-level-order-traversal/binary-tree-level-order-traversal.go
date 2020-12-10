package problem0102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 出现了新的 level
		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}

// dfs的常规写法
var res [][]int
func levelOrder1(root *TreeNode) [][]int {
	if root == nil{
		return nil
	}
	res = make([][]int, 0)
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int){
	if root == nil{
		return
	}
	if level == len(res){
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right,level+1)
}

// 广度优先

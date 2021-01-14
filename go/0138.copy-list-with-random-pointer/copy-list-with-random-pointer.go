package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	res := &Node{}
	p := res

	visited := make(map[*Node]*Node)

	// 先存next节点
	cur := head
	for cur != nil {
		p.Next = &Node{Val: cur.Val}
		visited[cur] = p.Next
		p = p.Next
		cur = cur.Next
	}

	cur = head
	p = res.Next

	// 获取random的深拷贝
	for cur != nil {
		p.Random = visited[cur.Random]
		p = p.Next
		cur = cur.Next
	}

	return res.Next
}

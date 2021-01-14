package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将问题转换成：将一个[]int转换成平衡二叉搜索树
func sortedListToBST(head *ListNode) *TreeNode {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return buildBST(arr, 0, len(arr)-1)
}

func buildBST(arr []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)>>1
	root := &TreeNode{Val: arr[mid]}

	root.Left = buildBST(arr, start, mid-1)
	root.Right = buildBST(arr, mid+1, end)
	return root
}

//快慢指针
// 快慢指针起初都指向头结点，分别一次走两步和一步，当快指针走到尾节点时，慢指针正好走到链表的中间。断成两个链表，分而治之。
// 为了断开，我们需要保存慢指针的前一个节点，因为单向链表的结点没有前驱指针。

func sortedListToBST0(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	return buildBST0(head, nil)
}

func buildBST0(head, tail *ListNode) *TreeNode {
	slow := head
	fast := head
	if head == tail {
		return nil
	}

	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}

	root := &TreeNode{Val: slow.Val}
	root.Left = buildBST0(head, slow)
	root.Right = buildBST0(slow.Next, tail)
	return root
}

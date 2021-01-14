package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/solution/java-ya-jie-dian-fei-di-gui-rong-yi-li-jie-yong-sh/
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy

	// 因为要对比cur.Next 与 cur.Next.Next 的Val是否相等，不能产生空指针
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			temp := cur.Next
			for temp != nil && temp.Next != nil && temp.Val == temp.Next.Val {
				temp = temp.Next
			}

			// temp.Next其实就是cur.Next.Next 不可能为nil
			cur.Next = temp.Next
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}

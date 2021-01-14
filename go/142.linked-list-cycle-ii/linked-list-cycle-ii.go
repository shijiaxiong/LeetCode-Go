package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针两次相遇
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	fast = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

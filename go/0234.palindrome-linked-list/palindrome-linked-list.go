package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度：O(N)
// 空间复杂度：O(N)
func isPalindrome(head *ListNode) bool {
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	start, end := 0, len(vals)
	for start < end {
		if vals[start] != vals[end] {
			return false
		}

		start++
		end--
	}

	return true
}

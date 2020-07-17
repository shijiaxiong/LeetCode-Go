package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	start, maxLength, length := 0, 0, 0
	for k, v := range []rune(s) {
		if i, ok := charIndex[v]; ok {
			start = i + 1
			length = k - start + 1
		} else {
			length += 1
		}

		charIndex[v] = k

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func main() {
	return
}

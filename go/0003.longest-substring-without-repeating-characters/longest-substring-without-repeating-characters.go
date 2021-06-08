package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	start, maxLength, length := 0, 0, 0
	for i, v := range []rune(s) {
		// 遍历过程中出现重复，重复位置如果在上次起始位置之后，就重新计算
		if lastIndex, ok := charIndex[v]; ok && lastIndex >= start {
			start = lastIndex + 1
			length = i - start + 1
		} else {
			length += 1
		}
		charIndex[v] = i
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func lengthOfLongestSubstring1(s string) int {
	start,end := 0,0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i],string(s[i]))
		if index==-1{
			if i+1>end{
				end=i+1
			}
		}else{
			start+=index+1
			end+=index+1
		}
	}
	return end-start
}


func main() {
	lengthOfLongestSubstring(`pwwkew`)
}

package main

import (
	"fmt"
)

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串

//func groupAnagrams(ss []string) [][]string {
//	tmp := make(map[int][]string, len(ss)/2)
//	for _, s := range ss {
//		c := encode(s)
//		tmp[c] = append(tmp[c], s)
//	}
//
//	res := make([][]string, 0, len(tmp))
//	for _, v := range tmp {
//		res = append(res, v)
//	}
//
//	return res
//}

// 方法一：排序数组分类
// 时间复杂度：
// 空间复杂度：
func groupAnagrams(strs []string) [][]string {
	tmp := make(map[string][]string, 0)

	for _, str := range strs {
		encodeResult := encode(str)
		tmp[encodeResult] = append(tmp[encodeResult], str)
	}

	res := make([][]string, 0)

	for _, v := range tmp {
		res = append(res, v)
	}

	return res
}

// 寻找一个唯一key的过程
func encode(str string) string {
	encodeSlice := make([]byte, 26)

	for _, char := range str {
		encodeSlice[char - 'a']++
	}

	return string(encodeSlice)
	//return *(*string)(unsafe.Pointer(&encodeSlice))
}

// 方法二：排序数组分类-质数法
func groupAnagrams2(strs []string) [][]string {
	prime := []int{5, 71, 37, 29, 2, 53, 59, 19, 11, 83, 79, 31, 43, 13, 7, 67, 97, 23, 17, 3, 41, 73, 47, 89, 61, 101}
	tmp := make(map[int][]string)

	for _, s := range strs {
		c := 1
		for i := range s {
			c *= prime[s[i]-'a']
		}

		tmp[c] = append(tmp[c], s)
	}

	res := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		res = append(res, v)
	}

	return res
}

func main() {
	param := []string{"eat","tea","tan","ate","nat","bat"}
	//param := []string{"eat", "tea", "ate"}

	fmt.Println(groupAnagrams(param))
}

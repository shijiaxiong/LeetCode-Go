package main

import (
	"fmt"
)

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
// 示例:
//
// 输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//
// 说明:
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
// Related Topics 字符串 回溯算法

var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

// 解法一：队列求解
// 时间复杂度：O(3^M×4^N)。M 是对应三个字母的数字个数，N 是对应四个字母的数字个数。
// 空间复杂度：O(3^M×4^N)。一共要生成3^M×4^N个结果。
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	ret := []string{""}

	for i := 0; i < len(digits); i++ {
		temp := []string{}

		for j := 0; j < len(ret); j++ {
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, ret[j]+m[digits[i]][k])
			}
		}

		ret = temp
	}

	return ret
}

// 解法二：回溯法求解
// 时间复杂度：O(3^M×4^N)。M 是对应三个字母的数字个数，N 是对应四个字母的数字个数。
// 空间复杂度：O(3^M×4^N)。一共要生成3^M×4^N个结果。
func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	ret := make([]string, 0)

	// 递归入口
	// 上一层产物-输入规模-递归层级-全局变量
	// offset递归层级可以通过上层产物去判断
	combination("", digits, 0, &ret)
	return ret
}

func combination(prefix string, digits string, offset int, ret *[]string) {
	if offset >= len(digits) {
		*ret = append(*ret, prefix)
		return
	}

	letters := m[digits[offset]]
	for i := 0; i < len(letters); i++ {
		combination(prefix + letters[i], digits, offset + 1, ret)
	}

}

// 回溯法求解 的另一种写法
func letterCombinations3(digits string) []string {
	if len(digits) == 0{
		return nil
	}

	res := make([]string,0)

	combination3(``, digits, &res)
	return res
}

func combination3(cur string, digits string, res *[]string){
	if len(cur) == len(digits) {
		*res = append(*res, cur)
		return
	}

	letters := m[digits[len(cur)]]

	for i:= 0; i<len(letters);i++{
		combination(cur +letters[i],digits,res)
	}

}


func main() {
	fmt.Println(letterCombinations(`23`))
	fmt.Println(letterCombinations2(`23`))
}

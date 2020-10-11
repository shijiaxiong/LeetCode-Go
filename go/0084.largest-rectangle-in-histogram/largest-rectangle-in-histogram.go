package main

import (
	"fmt"
)

// 单调栈 + 常数优化
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	fmt.Println(right)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i] - left[i] - 1) * heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func largestRectangleArea(heights []int) int {
	// 在 heights 中原本是非负的数字
	// 再首尾添加了 -2 和 -1 后，简化 for 循环中 begin 的计算
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)

	size := len(heights)

	// stack 中存的是 heights 的元素的索引号
	// stack 中索引号对应的 heights 中的值，是递增的。
	// e.g.
	//     stack = []int{1,3,5}，那么
	//     heights[1] < heights[3] < heights[5]
	stack := make([]int, 1, size)
	// 把 heights[0] 的索引号，已经放入了 stack
	// end 从 1 开始
	end := 1

	res := 0
	for end < size {
		// end 指向了新高，就把 end 放入 stack 后，指向下一个
		// stack中最后一个元素对应的高度 与 end 标识做对比
		if heights[stack[len(stack)-1]] < heights[end] {
			stack = append(stack, end)
			end++
			continue
		}

		begin := stack[len(stack)-2]
		index := stack[len(stack)-1]
		height := heights[index]
		fmt.Println(begin, index, height)
		// area 是 heights(begin:end) 之间的最大方块的面积，因为
		// anyone of heights(begin:index) > height > anyone of heights(index:end)
		area := (end - begin - 1) * height

		res = max(res, area)

		// h 的索引号已经没有必要留在 stack 中了
		stack = stack[:len(stack)-1]
	}

	return res
}

func largestRectangleArea2(heights []int) int {
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)

	size := len(heights)

	stack := make([]int, 1, size)
	end := 1
	res := 0

	for end < size {
		if heights[stack[len(stack)-1]] < heights[end] {
			stack = append(stack, end)
			end++
			continue
		}

		begin := stack[len(stack)-2]
		index := stack[len(stack)-1]
		height := heights[index]
		area := (end - begin - 1) * height
		res = max(res, area)
		stack = stack[:len(stack)-1]
	}
	return res
}


func main() {
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
}

package main

func bigger(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func smaller(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	length := len(height)
	if length <= 2 {
		return 0
	}

	left, right := make([]int, length), make([]int, length)

	left[0], right[length-1] = height[0], height[length-1]

	// left 与right 分别取左右的最大值
	for i := 1; i < length; i++ {
		left[i] = bigger(left[i-1], height[i])
		right[length-1-i] = bigger(right[length-i], height[length-1-i])
	}


	// water 是累计的每个i上可以存放的水量
	water := 0
	for i := 0; i < length; i++ {
		//fmt.Println(smaller(left[i], right[i]) - height[i])
		water += smaller(left[i], right[i]) - height[i]
	}

	return water
}

func main() {
	trap([]int{0,1,0,2,1,0,1,3,2,1,2,1})
}

// 单调栈
// https://leetcode-cn.com/problems/trapping-rain-water/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-8/
func trap1(height []int) int {
	sum := 0
	// 栈存储元素下标
	stack := []int{}
	current := 0
	for current < len(height) {
		// 保持栈内元素顺序递减，也就是当前元素不大于栈顶元素
		for len(stack) > 0 && height[current] > height[stack[len(stack) - 1]] {
			// 要取出的元素下标
			h := height[stack[len(stack) - 1]]

			// 出栈
			stack = stack[:len(stack) - 1]
			if len(stack) == 0 {
				break
			}

			// 两堵墙之间的距离
			distance := current - stack[len(stack) - 1] - 1
			m := smaller(height[current], height[stack[len(stack) - 1]])

			sum = sum + distance * (m - h)
		}

		stack = append(stack, current)
		current++
	}

	return sum
}

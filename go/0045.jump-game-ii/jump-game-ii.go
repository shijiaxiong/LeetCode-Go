package problem0045

func jump(nums []int) int {
	end := 0
	maxPosition := 0
	steps := 0

	for i:=0;i< len(nums)-1;i++{
		//找能跳的最远的
		maxPosition = max(maxPosition,nums[i] +i)
		//遇到边界，就更新边界，并且步数加一
		if i ==end {
			end= maxPosition
			steps++
		}
	}

	return steps
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}
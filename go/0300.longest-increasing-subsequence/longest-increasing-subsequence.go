package main

// 动态规划
// 状态转移方程：dp[i] = max(dp[i], dp[j]+1)
// 时间复杂度：O(N^2)
// 空间复杂度：O(N)
func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}

	dp := make([]int, length)
	for k := range dp {
		dp[k] = 1
	}

	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0

	for _, v := range dp {
		res = max(res, v)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 贪心+ 二分法
func lengthOfLIS1(nums []int) int {
	length := len(nums)

	if length < 2 {
		return length
	}

	//  长度为i+1的子序列末尾最小的是几
	tail := make([]int, length)
	tail[0] = nums[0]
	// end 表示有序数组 tail 的最后一个已经赋值元素的索引
	end := 0

	for i := 1;i < length;i++{
		if nums[i] > tail[end] {
			end++
			tail[end] = nums[i]
		} else {
			left := 0
			right := end

			for left < right {
				mid := left + (right - left)/2
				if tail[mid] < nums[i] {
					left = mid + 1
				} else {
					right = mid
				}
			}

			tail[left] = nums[i]
		}
	}

	end++
	// 输出队列
	// fmt.Println(tail[:end])
	return end
}

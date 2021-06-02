package problem0053

import (
	"fmt"
	"math"
)

// 模拟求和
func maxSubArray(nums []int) int {
	sum, maxSum := -1<<31, -1<<31
	for _, n := range nums {
		// sum+n < n，那就还不如直接从 n 开始统计。
		sum = max(sum+n, n)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划
//https://leetcode-cn.com/problems/maximum-subarray/solution/dong-tai-gui-hua-fen-zhi-fa-python-dai-ma-java-dai/
//dp[i] = max(nums[i], nums[i] + dp[i - 1])
func maxSubArray0(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}

		res = max(res, dp[i])
	}

	return res
}

// 动态规划压缩
// 降低空间复杂度
func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	pre := nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		pre = max(nums[i], nums[i]+pre)

		res = max(res, pre)
	}

	return res
}

// int result = INT_MIN;
//        int numsSize = int(nums.size());
//        int sum = 0;
//        for (int i = 0; i < numsSize; i++)
//        {
//            sum += nums[i];
//            result = max(result, sum);
//            //如果sum < 0，重新开始找子序串
//            if (sum < 0)
//            {
//                sum = 0;
//            }
//        }
//
//        return result;

// 贪心算法
func maxSubArray2(nums []int) int {
	max := math.MinInt32
	sum := 0

	for i := 0; i < len(nums); i++ {

		if sum > 0 {
			sum = sum + nums[i]
		} else {
			sum = nums[i]
		}

		if sum > max {
			max = sum
		}

	}

	return max
}

// 贪心算法返回最大子序和的最大值
// https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-fan-hui-zi-shu-zu-qu-jian-by-mitty/
func maxSubArray3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 子序列和最大值
	max := math.MinInt32
	// 某段子序列和
	sum := 0

	// 最大子序列和的左右下标
	start, end := 0, 0
	// 某段子序列的左右下标
	subStart, subEnd := 0, 0

	for i := 0; i < len(nums); i++ {
		//	sum > 0时，sum+=num[i]
		//  sum < 0，就没有必要再加了，因为下一个数，如果是负的，你加上负的，只会更小，
		//  如果是正数，你加上他，也不会比直接等于它大，sum < 0,直接让 sum 等于当前的元素
		if sum > 0 {
			sum += nums[i]
			subEnd++
		} else {
			sum = nums[i]
			subStart = i
			subEnd = i
		}

		if sum > max {
			max = sum
			start = subStart
			end = subEnd
		}
	}

	fmt.Println(start, end)
	return max
}

func maxSubArray4(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return maxSubArraySum(nums, 0, len(nums)-1)
}

func maxCrossingSum(nums []int, left, mid, right int) int {
	// 一定会包含 nums[mid] 这个元素
	sum := 0
	leftSum := math.MinInt32

	// 左半边包含 nums[mid] 元素，最多可以到什么地方
   // 走到最边界，看看最值是什么
   // 计算以 mid 结尾的最大的子数组的和
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	sum = 0
	rightSum := math.MinInt32
	// 右半边不包含 nums[mid] 元素，最多可以到什么地方
	// 计算以 mid+1 开始的最大的子数组的和
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}

func maxSubArraySum(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) >> 1

	return max3(
		maxSubArraySum(nums, left, mid),
		maxSubArraySum(nums, mid+1, right),
		maxCrossingSum(nums, left, mid, right))
}

func max3(a, b, c int) int {
	max := 0
	if a > b {
		max = a
	} else {
		max = b
	}

	if max > c {
		return max
	}

	return c
}

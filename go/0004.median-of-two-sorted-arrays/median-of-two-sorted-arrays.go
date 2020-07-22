package main

import (
	"fmt"
)

//先将两个数组合并，然后找中位数
//时间复杂度：遍历全部数组 (m+n)
//空间复杂度：开辟了一个数组，保存合并后的两个数组 O(m+n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, i := len(nums1), 0
	l2, j := len(nums2), 0

	if l1 == 0 {
		if l2%2 == 0 {
			return float64(nums2[l2/2-1]+nums2[l2/2]) / 2.0
		} else {
			return float64(nums2[l2/2])
		}
	}

	if l2 == 0 {
		if l1%2 == 0 {
			return float64(nums1[l1/2-1]+nums1[l1/2]) / 2.0
		} else {
			return float64(nums1[l1/2])
		}
	}

	res := make([]int, l1+l2)

	for k := 0; k < l1+l2; k++ {
		if i == l1 && j != l2 {
			res[k] = nums2[j]
			j++
			continue
		}

		if j == l2 && i != l1 {
			res[k] = nums1[i]
			i++
			continue
		}

		if nums1[i] < nums2[j] {
			res[k] = nums1[i]
			i++

		} else {
			res[k] = nums2[j]
			j++
		}
	}

	count := l1 + l2

	if count%2 == 0 {
		return float64(res[count/2-1]+res[count/2]) / 2.0
	} else {
		return float64(res[count/2])
	}

}

func main() {
	res := findMedianSortedArrays([]int{ 2}, []int{})
	fmt.Println(res)
}

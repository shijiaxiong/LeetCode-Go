package main

import (
	"fmt"
)

//给你一个整数数组 nums，请你将该数组升序排列。
//
//
//
//
//
//
// 示例 1：
//
// 输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
//
//
// 示例 2：
//
// 输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 50000
// -50000 <= nums[i] <= 50000
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	quickSort(nums)
	return nums
}

// 方法一： 快排
func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	reference := nums[0]
	head, tail := 0, len(nums)-1
	i := 1

	for head < tail {
		if nums[i] < reference {
			nums[head], nums[i] = nums[i], nums[head]
			head++
			i++
		} else {
			nums[tail], nums[i] = nums[i], nums[tail]
			tail--
		}
	}

	quickSort(nums[:head])
	quickSort(nums[head+1:])
	return
}

func quickSort1(nums []int, left, right int) {
	if left >= right {
		return
	}

	p := partitionL(nums, left, right)
	quickSort1(nums, left, p-1)
	quickSort1(nums, p+1, right)
}

// 以右边数字为基准点
func partition(nums []int, lo, hi int) int {
	// 找到基准点
	pivot := nums[hi]

	i := lo - 1
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			i++
			nums[j], nums[i] = nums[i], nums[j]
			fmt.Println(nums, i)
		}
	}
	nums[i+1], nums[hi] = nums[hi], nums[i+1]
	return i + 1
}

//从左往右依次是low high unvisited部分
func partitionL(nums []int, left, right int) int {

	//rand.Seed(time.Now().Unix())
	nums[left], nums[(left+right)>>1] = nums[(left+right)>>1], nums[left]

	// 以最左边的数字为基准
	pivot := nums[left]

	i := left

	for j := left + 1; j <= right; j++ {
		if nums[j] < pivot {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}

	nums[i], nums[left] = nums[left], nums[i]
	return i
}

// 双指针双向移动
func partition3(nums []int, left, right int) int {
	i := left
	j := right
	nums[i], nums[(i+j)>>1] = nums[(i+j)>>1], nums[i]
	pivot := nums[i]

	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		// 从右往左找到了第一个小于 基准值的值
		nums[i] = nums[j]

		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}

	nums[i] = pivot

	return i
}

// 归并排序
// https://leetcode-cn.com/problems/sort-an-array/solution/golang-by-xilepeng-2/
func mergeSort(nums []int, left, right int) {
	if left < right {
		mid := (left + right) >> 1
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {
	i := left
	j := mid + 1
	tmp := []int{}

	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}

	if i <= mid {
		tmp = append(tmp, nums[i:mid+1]...)
	} else {
		tmp = append(tmp, nums[j:right+1]...)
	}

	copy(nums[left:right+1], tmp)
	// for k := 0; k < len(tmp); k++ {
	// 	nums[left+k] = tmp[k]
	// }
}

// 堆排序
func heapSort(nums []int) []int {
	lens := len(nums)
	// 建堆
	// 只需要从 i = (len - 1) / 2 这个位置开始逐层下移
	for i := lens/2 - 1; i >= 0; i-- {
		heapify(nums, lens, i)
	}

	// 排序
	// 循环不变量：区间 [0, i] 堆有序
	//逐渐减少堆有序的部分
	for i := lens - 1; i > 0; i-- {
		// 把堆顶元素（当前最大）交换到数组末尾
		nums[0], nums[i] = nums[i], nums[0]

		// 下标 0 位置下沉操作，使得区间 [0, i] 堆有序
		heapify(nums, i, 0)
	}

	return nums
}

// 维护一个小堆,思路是找到小堆中的最大值，下标为largest,
// n 数组长度
// i 待维护节点的下标
func heapify(nums []int, n, i int) {
	// 当前节点下标
	largest := i
	// 左节点下标
	lson := i*2 + 1
	// 右节点下标
	rson := i*2 + 2

	// 左节点下标在有效范围内,左节点大于最大值。那么当前最大值在左子树
	if lson < n && nums[largest] < nums[lson] {
		largest = lson
	}
	if rson < n && nums[largest] < nums[rson] {
		largest = rson
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		// 这里三元小堆中的最大值已经换到了i下标
		// largest下标是函数入口处i 下标的值
		heapify(nums, n, largest)
	}
}

func main() {
	nums := []int{9, 5, 2, 11, 12, 4, 3, 1, 7}

	fmt.Println(insertSort(nums))
}

// 冒泡排序
func bubbleSort(nums []int) []int {
	// 第一层循环是元素的个数
	for i := 0; i < len(nums); i++ {
		// 先默认数组是有序的，只要发生一次交换，就必须进行下一轮比较，
		// 如果在内层循环中，都没有执行一次交换操作，说明此时数组已经是升序数组
		sorted := true

		// 第二层循环是比较的次数
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				sorted = false
			}
		}

		if sorted {
			break
		}
	}

	return nums
}

// 选择排序
func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minIndex := i

		for j := i + 1; j < len(nums); j++ {
			// 找到更小的元素和下标
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

	return nums
}

// 插入排序
func insertSort(nums []int) []int {

	for i := 1; i < len(nums); i++ {
		// 先暂存这个元素，然后之前元素逐个后移，留出空位
		temp := nums[i]

		j := i

		for j > 0 && nums[j - 1] > temp {
			nums[j] = nums[j - 1]
			j--
		}
		nums[j] = temp
	}

	return nums
}

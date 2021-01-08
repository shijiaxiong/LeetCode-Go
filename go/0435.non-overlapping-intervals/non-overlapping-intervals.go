package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 按右边界排序
	// 如果使用左边界排序， 需要在if判断处增加对preEnd 与 curEnd的比较，选取较小值
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	fmt.Println(intervals)
	// 记录区间尾部位置
	end := intervals[0][1]
	// 需要移除的数量
	var res int
	for i := 1; i < len(intervals); i++ {
		// 前一个的结束时间 大于 当前的开始时间就会产生重叠
		if end > intervals[i][0] {
			// 如果出现重叠区间 移除一个 所以count+1
			res++
			continue
		}
		end = intervals[i][1]
	}

	return res
}

func main() {
	eraseOverlapIntervals([][]int{{1, 2}, {1, 3}, {2, 3}, {3, 4}, {1, 3}})
}
//[[1 2] [2 3] [1 3] [1 3] [3 4]]
// [[1,100],[11,22],[1,11],[2,12]]  ans = 2

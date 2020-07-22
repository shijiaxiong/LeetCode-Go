package main

import (
	"fmt"
)

//class Solution {
//    public int maxArea(int[] height) {
//        int i = 0, j = height.length - 1, res = 0;
//        while(i < j){
//            res = height[i] < height[j] ?
//                Math.max(res, (j - i) * height[i++]):
//                Math.max(res, (j - i) * height[j--]);
//        }
//        return res;
//    }
//}
//

func maxArea(height []int) int {
	i, j, max := 0, len(height)-1, 0

	for i < j {
		if height[i] < height[j] {
			max = getMax(max, (j-i) * height[i])
			i++
		} else {
			max = getMax(max, (j-i) * height[j])
			j--
		}
	}

	return max
}

func getMax(i, j int) int {
	if i <= j {
		return j
	}
	return i
}


func maxArea1(height []int) int {
	i, j, max := 0, len(height) - 1, 0

	for i < j {
		a,b := height[i], height[j]
		h := 0
		if a >= b {
			h = b
		} else {
			h = a
		}

		area := (j - i) * h
		if area > max {
			max = area
		}

		if a < b {
			i++
		} else {
			j--
		}
	}

	return max
}

func main() {
	res := maxArea([]int{1,8,6,2,5,4,8,3,7})
	fmt.Println(res)
}

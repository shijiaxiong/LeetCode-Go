package main

import (
	"math"
)


func reverse(x int) int {
	res :=0
	for x!=0{
		temp :=x%10

		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && temp>7) {
			return 0
		}
		
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && temp < -8) {
			return 0
		}

		res = res *10 +temp
		x /=10
	}

	return res
}

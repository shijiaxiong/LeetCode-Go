package proble0069

// 牛顿法
func mySqrt0(x int) int {
	a :=x
	for a*a >x {
		a= (a + x/a)/2
	}
	return a
}

// 二分法
func mySqrt(x int) int {
	l, h:= 1, x/2 + 1
	for l <= h {
		mid := l + (h - l) / 2
		if mid == x / mid {
			return mid
		} else if mid > x / mid {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return h
}
package main

//按照题目的描述，可以总结如下规则：
//
//罗马数字由 I,V,X,L,C,D,M 构成；
//当小值在大值的左边，则减小值，如 IV=5-1=4；
//当小值在大值的右边，则加小值，如 VI=5+1=6；
//由上可知，右值永远为正，因此最后一位必然为正。
//一言蔽之，把一个小值放在大值的左边，就是做减法，否则为加法。

func romanToInt(s string) int {
	sum := 0
	preNum := getValue(s[0])
	for i := 1; i < len(s); i++ {
		num := getValue(s[i])

		if preNum < num {
			sum -= preNum
		} else {
			sum += preNum
		}

		preNum = num
	}

	//最后一位直接加
	sum = preNum + sum

	return sum
}

func getValue(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}

}

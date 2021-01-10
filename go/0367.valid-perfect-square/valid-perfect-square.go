package main

func isPerfectSquare(num int) bool {
	num1 := 1
	for num > 0 {
		num -= num1
		num1 += 2
	}

	return num == 0
}

package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}

	digits = append([]int{1}, digits...)
	return digits
}

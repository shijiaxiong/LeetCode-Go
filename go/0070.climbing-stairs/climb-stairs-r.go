package problem0070

func climbStairsT(n int) int {
	mem := make(map[int]int)
	return recursive(0, n, mem)
}

func recursive(i, n int, mem map[int]int) int {
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	if mem[i] > 0 {
		return mem[i]
	}

	return recursive(i+1, n, mem) + recursive(i+2, n, mem)
}

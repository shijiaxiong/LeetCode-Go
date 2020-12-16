package problem0122

//贪心算法
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	res := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

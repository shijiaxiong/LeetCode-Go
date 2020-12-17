package problem0714

// 贪心算法
func maxProfit0(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}

	ans := 0
	inHand := prices[0]

	for i := 1; i < len(prices); i++ {
		// 找谷
		if prices[i] < inHand {
			inHand = prices[i]
		} else if prices[i]-fee > inHand { //找减去手续费的峰值
			ans += prices[i] - fee - inHand
			inHand = prices[i] - fee
		}
	}

	return ans
}



package main

// 动态规划
// dp
// 存储中间态：dp[i]为s[0...i]的译码总数
// 递推公式：
// 当s[i] == '0', 若s[i-1]=='1'||'2', 则 dp[i] = dp[i-2]
// 当s[i-1] == '1', dp[i] = dp[i-1] + dp[i-2] //相当于 跨一步 + 跨两步
// 当s[i-1] == '2', 若s[i]<='6', dp[i] = dp[i-1] + dp[i-2], 否则dp[i]=dp[i-1]

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	pre, curr := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := curr
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || s[i-1] == '2' && s[i] >= '1' && s[i] <= '6' {
			curr = curr + pre
		}
		pre = tmp
	}

	return curr
}

// 动态规划
func numDecodings0(s string) int {
	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	// dp[0]
	dp[0],dp[1] = 1,1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dp[i+1] = dp[i-1]
			} else {
				return 0
			}
		} else if s[i-1] == '1' || s[i-1] == '2' && s[i] >= '1' && s[i] <= '6' {
			dp[i+1] = dp[i] + dp[i-1]
		}else{
			dp[i+1] = dp[i]
		}
	}

	return dp[len(s)]
}

func main() {
	numDecodings0("120")
}

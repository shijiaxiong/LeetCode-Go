package main

// class Solution {
//    public int nthUglyNumber(int n) {
//        int a = 0, b = 0, c = 0;
//        int[] dp = new int[n];
//        dp[0] = 1;
//        for(int i = 1; i < n; i++) {
//            int n2 = dp[a] * 2, n3 = dp[b] * 3, n5 = dp[c] * 5;
//            dp[i] = Math.min(Math.min(n2, n3), n5);
//            if(dp[i] == n2) a++;
//            if(dp[i] == n3) b++;
//            if(dp[i] == n5) c++;
//        }
//        return dp[n - 1];
//    }
//}
//
// 动态规划
func nthUglyNumber(n int) int {
	a := 0
	b := 0
	c := 0

	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		n2 := dp[a] * 2
		n3 := dp[b] * 3
		n5 := dp[c] * 5
		dp[i] = min(min(n2, n3), n5)

		if dp[i] == n2 {
			a++
		}

		if dp[i] == n3 {
			b++
		}

		if dp[i] == n5 {
			c++
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

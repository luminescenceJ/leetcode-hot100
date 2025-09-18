package repo

import "math"

func numSquares(n int) int {
	// dp[i] 表示和为i的最少平方数量
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := range dp {
		for j := 0; j*j <= i; j++ {
			dp[i] = min(dp[i-j*j]+1, dp[i])
		}
	}
	return dp[n]
}

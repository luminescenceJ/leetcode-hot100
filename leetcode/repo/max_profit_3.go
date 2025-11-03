package repo

func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	if len(prices) == 2 {
		return max(prices[1]-prices[0], 0)
	}

	// dp[i][2] // 截止到i天当天是否持有股票的最多赚钱
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	// 0表示无股票无冷却期 1表示有股票 2表示无股票有冷却期
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][2])
}

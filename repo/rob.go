package repo

func rob(nums []int) int {
	// dp[i][2]
	// dp[i][0]表示不偷第i家的最大金额
	// dp[i][1]表示偷的最大金额
	if len(nums) == 0 {
		return 0
	}
	dp := make([][2]int, len(nums))
	for i := range dp {
		dp[i] = [2]int{}
	}
	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

package repo

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	// sum-2*neg = target
	neg := (sum - target) / 2
	// 从nums找到和为neg的不连续不重复序列的数量
	if (sum-target) < 0 || (sum-target)%2 != 0 {
		return 0
	}
	// dp[i][amount]表示寻找到第i个数，和为amount的数量
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i := range nums {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= nums[i] {
				dp[i+1][j] += dp[i][j-nums[i]]
			}
		}
	}
	return dp[len(nums)][neg]
}

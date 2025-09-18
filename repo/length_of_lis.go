package repo

func lengthOfLIS(nums []int) int {
	// dp[i]表示以nums[i]结尾的最长严格递增子序列长度
	dp := make([]int, len(nums))
	var res int
	for i := range dp {
		dp[i] = 1
		if i == 0 {
			res = max(res, dp[i])
			continue
		}
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			res = max(res, dp[i])
		}
	}
	return res
}

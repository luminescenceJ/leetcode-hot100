package repo

func maxProduct(nums []int) int {
	// dp[i]表示到nums[i]的最大连续乘积子数组
	// mindp[i]表示nums[i]的最小连续乘积子数组
	var res int = nums[0]
	dp := make([]int, len(nums))
	mindp := make([]int, len(nums))
	dp[0], mindp[0] = nums[0], nums[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = max(max(dp[i-1]*nums[i], nums[i]), mindp[i-1]*nums[i])
		mindp[i] = min(nums[i], min(dp[i-1]*nums[i], mindp[i-1]*nums[i]))
		res = max(res, dp[i])
	}
	return res
}

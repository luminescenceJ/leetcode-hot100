package repo

func numTrees(n int) int {
	// n个数字顺序排列
	// 枚举每个数字作为根节点有多少种二叉搜索树
	dp := make([]int, 20)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		cur := 0
		for j := 0; j < i; j++ {
			cur += dp[i-j-1] * dp[j]
		}
		dp[i] = cur
	}
	return dp[n]
}

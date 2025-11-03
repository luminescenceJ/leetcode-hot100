package repo

func minPathSum(grid [][]int) int {
	// dp[i][j]表示到达grid[i][j]的最小路径
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := range grid {
		if i == 0 {
			continue
		}
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := range grid[0] {
		if i == 0 {
			continue
		}
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

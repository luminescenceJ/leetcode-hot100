package repo

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxArea := 0
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0])) // dp[i][j]表示以i,j为右下角坐标的正方形的边长
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
					maxArea = max(maxArea, dp[i][j])
					continue
				}
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			maxArea = max(maxArea, dp[i][j])
		}
	}
	return maxArea * maxArea
}

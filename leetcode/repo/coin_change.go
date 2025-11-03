package repo

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	// dp[rest] 表示凑齐rest个硬币所需最少数量
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	sort.Ints(coins)
	// 硬币升序排列
	for i := range dp {
		for _, c := range coins {
			if c > i {
				break
			}
			if dp[i-c] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

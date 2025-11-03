package repo

func wordBreak(s string, wordDict []string) bool {
	// dp[i]表示s[i]截止可以被表示
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, w := range wordDict {
			if len(w) <= i {
				if s[i-len(w):i] == w && dp[i-len(w)] {
					dp[i] = true
				}
			}
		}
	}
	return dp[len(s)]
}

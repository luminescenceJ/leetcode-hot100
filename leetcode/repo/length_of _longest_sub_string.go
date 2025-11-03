package repo

func lengthOfLongestSubstring(s string) int {
	// 滑动窗口记录内部字符数量
	counter := make(map[byte]int)
	var l, r = 0, 0
	var res int
	for r < len(s) {
		for counter[s[r]-'a'] > 0 {
			counter[s[l]-'a']--
			l++
		}
		counter[s[r]-'a']++
		r++
		res = max(res, r-l)
	}
	return res
}

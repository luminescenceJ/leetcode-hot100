package repo

func countSubstrings(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		if i != 0 {
			res += diffusion(i-1, i, s)
		}
		res += diffusion(i, i, s)
	}
	return res
}

func diffusion(l, r int, s string) int {
	// 从l,r位置向两边扩散的最多回文子串数量
	var res int
	for l >= 0 && r <= len(s)-1 && l <= r {
		if s[l] != s[r] {
			return res
		}
		res += 1
		l--
		r++
	}
	return res
}

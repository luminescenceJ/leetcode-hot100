package repo

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	flag := [26]int{}
	for _, str := range p {
		flag[str-'a']++
	}

	pattern := [26]int{}
	for i := 0; i < len(p)-1; i++ {
		pattern[s[i]-'a']++
	}
	res := []int{}
	for i := len(p) - 1; i < len(s); i++ {
		pattern[s[i]-'a']++
		if pattern == flag {
			res = append(res, i-len(p)+1)
		}
		pattern[s[i-len(p)+1]-'a']--
	}
	return res
}

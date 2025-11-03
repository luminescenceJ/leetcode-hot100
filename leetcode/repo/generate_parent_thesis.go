package repo

func generateParenthesis(n int) []string {
	var res []string
	cur := []byte{}
	var dfs func(int, int)
	dfs = func(l, r int) {
		if l < r || l > n || r > n {
			return
		}
		if l == n && r == n {
			res = append(res, string(cur))
			return
		}

		cur = append(cur, '(')
		dfs(l+1, r)
		cur = cur[:len(cur)-1]

		cur = append(cur, ')')
		dfs(l, r+1)
		cur = cur[:len(cur)-1]

	}
	dfs(0, 0)
	return res
}
